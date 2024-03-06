package utilities

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
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

var Scheduler *cron.Cron

func schedulerTick() {
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

	Scheduler = cron.New()
	_, schedulerError := Scheduler.AddFunc(
		"* * * * *",
		schedulerTick,
	)
	if schedulerError != nil {
		log.Fatal(schedulerError)
	}
	Scheduler.Start()
	log.Println("CRON started")
}
