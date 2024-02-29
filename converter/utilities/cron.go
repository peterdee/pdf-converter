package utilities

import (
	"context"
	"fmt"
	"log"
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

	session, sessionError := database.Client.StartSession()
	if sessionError != nil {
		log.Fatal(sessionError)
	}
	defer session.EndSession(ctx)

	_, transactionError := session.WithTransaction(
		ctx,
		func(sctx mongo.SessionContext) (interface{}, error) {
			var entry database.QueueEntry
			queryError := database.Queue.FindOne(
				sctx,
				bson.D{
					{Key: "status", Value: constants.QUEUE_STATUSES.InProgress},
				},
			).Decode(&entry)
			if queryError != nil && queryError != mongo.ErrNoDocuments {
				log.Fatal(queryError)
			}
			if queryError == mongo.ErrNoDocuments {
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
					return new(interface{}), nil
				}

				filePath := fmt.Sprintf("./processing/%s/%s.pdf", queueItem.UID, queueItem.UID)
				resultMask := fmt.Sprintf("./processing/%s/p.jpg", queueItem.UID)
				result, convertError := imagick.ConvertImageCommand([]string{
					"-density 150",
					filePath,
					resultMask,
				})
				fmt.Println(result, convertError)

				zipError := CreateZipArchive(queueItem.UID)
				if zipError != nil {
					log.Fatal(zipError)
				}
			}
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
	// Scheduler.Stop()
	// Scheduler.Start()
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
