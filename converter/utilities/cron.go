package utilities

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/julyskies/gohelpers"
	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/gographics/imagick.v3/imagick"

	"converter/constants"
	"converter/database"
)

func archiveRemovalTick() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	archiveStoredDays := constants.ARCHIVE_STORED_DAYS
	if value := os.Getenv(constants.ENV_NAMES.ArchiveStoredDays); value != "" {
		parsed, parsingError := strconv.Atoi(value)
		if parsingError == nil {
			archiveStoredDays = parsed
		}
	}

	var deletableEntries []database.QueueEntry
	cursor, queryError := database.Queue.Find(
		ctx,
		bson.D{
			{Key: "downloadedAt", Value: bson.D{
				{
					Key:   "$lt",
					Value: gohelpers.MakeTimestamp() - (60 * 60 * 24 * int64(archiveStoredDays)),
				},
			}},
		},
	)
	if queryError != nil && queryError != mongo.ErrNoDocuments {
		log.Fatal(queryError)
	}
	if queryError == mongo.ErrNoDocuments {
		return
	}

	if decodeError := cursor.All(ctx, &deletableEntries); decodeError != nil {
		log.Fatal(decodeError)
	}
	if len(deletableEntries) == 0 {
		return
	}

	for _, entry := range deletableEntries {
		directoryPath := fmt.Sprintf("./processing/%s", entry.UID)
		os.RemoveAll(directoryPath)
	}

	dirContent, dirError := os.ReadDir("./processing")
	if dirError != nil {
		log.Fatal(dirError)
	}
	for _, entry := range dirContent {
		var databaseEntry database.QueueEntry
		entryError := database.Queue.FindOne(
			ctx,
			bson.D{{
				Key:   "uid",
				Value: entry.Name(),
			}},
		).Decode(&databaseEntry)
		if entryError == mongo.ErrNoDocuments {
			os.RemoveAll(fmt.Sprintf("./processing/%s", entry.Name()))
		}
	}
}

func processingTick() {
	initialContext, cancelInitialContext := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancelInitialContext()

	var inProgress database.QueueEntry
	queryError := database.Queue.FindOne(
		initialContext,
		bson.D{
			{Key: "status", Value: constants.QUEUE_STATUSES.InProgress},
		},
	).Decode(&inProgress)
	if queryError != nil && queryError != mongo.ErrNoDocuments {
		log.Fatal(queryError)
	}
	if inProgress.UID != "" {
		return
	}

	var queueItem database.QueueEntry
	queryError = database.Queue.FindOne(
		initialContext,
		bson.D{
			{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
		},
		options.FindOne().SetSort(bson.D{{Key: "createdAt", Value: 1}}),
	).Decode(&queueItem)
	if queryError != nil && queryError != mongo.ErrNoDocuments {
		log.Fatal(queryError)
	}
	if queryError == mongo.ErrNoDocuments {
		return
	}

	filePath := fmt.Sprintf("./processing/%s/%s.pdf", queueItem.UID, queueItem.UID)
	if _, existsError := os.Stat(filePath); errors.Is(existsError, os.ErrNotExist) {
		database.Queue.DeleteOne(initialContext, bson.D{{Key: "uid", Value: queueItem.UID}})
		return
	}
	_, queryError = database.Queue.UpdateOne(
		initialContext,
		bson.D{{Key: "uid", Value: queueItem.UID}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "status", Value: constants.QUEUE_STATUSES.InProgress},
				{Key: "updatedAt", Value: gohelpers.MakeTimestamp()},
			}},
		},
	)
	if queryError != nil {
		log.Fatal(queryError)
	}

	resultMask := fmt.Sprintf("./processing/%s/p.jpg", queueItem.UID)
	_, convertError := imagick.ConvertImageCommand([]string{
		"-density 150",
		filePath,
		resultMask,
	})
	if convertError != nil {
		log.Fatal(convertError)
	}
	if zipError := CreateZipArchive(queueItem.UID); zipError != nil {
		log.Fatal(zipError)
	}
	if deleteError := DeleteFiles(queueItem.UID); deleteError != nil {
		log.Fatal(deleteError)
	}

	postContext, cancelPostContext := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancelPostContext()
	_, queryError = database.Queue.UpdateOne(
		postContext,
		bson.D{{Key: "uid", Value: queueItem.UID}},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "status", Value: constants.QUEUE_STATUSES.Processed},
				{Key: "updatedAt", Value: gohelpers.MakeTimestamp()},
			}},
		},
	)
	if queryError != nil {
		log.Fatal(queryError)
	}
	log.Printf("Processed UID %s", queueItem.UID)
}

func LaunchCRON() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	_, updateError := database.Queue.UpdateMany(
		ctx,
		bson.D{
			{Key: "status", Value: constants.QUEUE_STATUSES.InProgress},
		},
		bson.D{
			{Key: "$set", Value: bson.D{
				{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
				{Key: "updatedAt", Value: gohelpers.MakeTimestamp()},
			}},
		},
	)
	if updateError != nil {
		log.Fatal(updateError)
	}

	ArchiveRemovalScheduler := cron.New()
	_, schedulerError := ArchiveRemovalScheduler.AddFunc(
		"* * * * *", // TODO: proper schedule
		archiveRemovalTick,
	)
	if schedulerError != nil {
		log.Fatal(schedulerError)
	}
	ArchiveRemovalScheduler.Start()

	ProcessingScheduler := cron.New()
	_, schedulerError = ProcessingScheduler.AddFunc(
		"* * * * *",
		processingTick,
	)
	if schedulerError != nil {
		log.Fatal(schedulerError)
	}
	ProcessingScheduler.Start()

	log.Println("CRON started")
}
