package utilities

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"gopkg.in/gographics/imagick.v3/imagick"

	"converter/constants"
	"converter/database"
)

var Scheduler *cron.Cron

func schedulerTick() {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// TODO: refactoring

	session, sessionError := database.Client.StartSession()
	if sessionError != nil {
		log.Fatal(sessionError)
	}
	defer session.EndSession(ctx)

	var converting bool = false
	var selectedUID string
	_, transactionError := session.WithTransaction(
		ctx,
		func(sctx mongo.SessionContext) (interface{}, error) {
			var inProgress database.QueueEntry
			queryError := database.Queue.FindOne(
				sctx,
				bson.D{
					{Key: "status", Value: constants.QUEUE_STATUSES.InProgress},
				},
			).Decode(&inProgress)
			if queryError != nil && queryError != mongo.ErrNoDocuments {
				log.Fatal(queryError)
			}
			if inProgress.ID.Hex() != "" {
				converting = true
				return new(interface{}), nil
			}

			var queueItem database.QueueEntry
			getDocumentError := database.Queue.FindOne(
				sctx,
				bson.D{},
				options.FindOne().SetSort(bson.D{{Key: "createdAt", Value: 1}}),
			).Decode(&queueItem)
			if getDocumentError != nil && getDocumentError != mongo.ErrNoDocuments {
				log.Fatal(getDocumentError)
			}
			if getDocumentError == mongo.ErrNoDocuments {
				converting = true
				return new(interface{}), nil
			}

			filePath := fmt.Sprintf("./processing/%s/%s.pdf", queueItem.UID, queueItem.UID)
			if _, existsError := os.Stat(filePath); errors.Is(existsError, os.ErrNotExist) {
				log.Println("File not found", filePath)
				database.Queue.DeleteOne(sctx, bson.D{{Key: "uid", Value: queueItem.UID}})
				return new(interface{}), nil
			}
			selectedUID = queueItem.UID
			fmt.Println(queueItem)
			updateResult, updateError := database.Queue.UpdateOne(
				sctx,
				bson.D{{Key: "_id", Value: queueItem.UID}},
				bson.D{
					{Key: "$set", Value: bson.D{
						{Key: "status", Value: constants.QUEUE_STATUSES.InProgress},
					}},
				},
			)
			fmt.Println(updateResult, updateError)
			return new(interface{}), nil
		},
		options.Transaction().SetWriteConcern(writeconcern.Majority()),
	)
	if transactionError != nil {
		log.Fatalf(
			"%s: %v",
			constants.ERROR_MESSAGES.MongoTransactionError,
			transactionError,
		)
	}

	if !converting {
		filePath := fmt.Sprintf("./processing/%s/%s.pdf", selectedUID, selectedUID)
		resultMask := fmt.Sprintf("./processing/%s/p.jpg", selectedUID)
		_, convertError := imagick.ConvertImageCommand([]string{
			"-density 150",
			filePath,
			resultMask,
		})
		if convertError != nil {
			log.Fatal(convertError)
		}
		zipError := CreateZipArchive(selectedUID)
		if zipError != nil {
			log.Fatal(zipError)
		}
		deleteError := DeleteFiles(selectedUID)
		if deleteError != nil {
			log.Fatal(deleteError)
		}
		// TODO: finalize
	}
}

func LaunchCRON() {
	Scheduler = cron.New()
	_, schedulerError := Scheduler.AddFunc(
		"* * * * *",
		func() {
			schedulerTick()
		},
	)
	if schedulerError != nil {
		log.Fatal(schedulerError)
	}
	Scheduler.Start()
	log.Println("CRON started")
}
