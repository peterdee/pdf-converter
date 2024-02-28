package utilities

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/robfig/cron/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"splitter/constants"
	"splitter/database"
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
				// TODO: at this point processing should be started
				fmt.Println("start processing")
				// imagick.ConvertImageCommand([]string{
				// 	"-density",s
				// 	"150",
				// 	"/presentation.pdf",
				// 	"-quality",
				// 	"90",
				// 	"test.jpg",
				// })
			}
			return new(interface{}), nil
		},
	)
	if transactionError != nil {
		log.Fatal(transactionError)
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
