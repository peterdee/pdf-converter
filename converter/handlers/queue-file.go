package handlers

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/julyskies/gohelpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"converter/constants"
	"converter/database"
	"converter/utilities"
)

func QueueFile(encoded string, filename string) (*QueueFileResult, error) {
	hash := utilities.CreateMD5Hash(fmt.Sprintf(
		"%s:%s:%d",
		encoded,
		filename,
		gohelpers.MakeTimestamp(),
	))
	dirPath := fmt.Sprintf("./processing/%s", hash)
	mkdirError := os.Mkdir(dirPath, 0777)
	if mkdirError != nil {
		return nil, mkdirError
	}

	bytes, decodeError := hex.DecodeString(encoded)
	if decodeError != nil {
		return nil, decodeError
	}
	os.WriteFile(fmt.Sprintf("%s/%s.pdf", dirPath, hash), bytes, 0777)

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	session, sessionError := database.Client.StartSession()
	if sessionError != nil {
		return nil, sessionError
	}
	defer session.EndSession(ctx)

	result, transactionError := session.WithTransaction(
		ctx,
		func(tctx mongo.SessionContext) (interface{}, error) {
			now := gohelpers.MakeTimestamp()
			_, insertionError := database.Queue.InsertOne(
				tctx,
				bson.D{
					{Key: "createdAt", Value: now},
					{Key: "downloadedAt", Value: now},
					{Key: "originalFileName", Value: filename},
					{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
					{Key: "totalDownloads", Value: 0},
					{Key: "uid", Value: hash},
					{Key: "updatedAt", Value: now},
				},
			)
			if insertionError != nil {
				return nil, insertionError
			}

			count, countError := database.Queue.CountDocuments(
				tctx,
				bson.D{
					{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
				},
			)
			if countError != nil {
				return nil, countError
			}

			var queueResult = &QueueFileResult{
				Filename:    filename,
				QueuedItems: count,
				Status:      constants.QUEUE_STATUSES.Queued,
				UID:         hash,
			}
			return queueResult, nil
		},
		options.Transaction().SetWriteConcern(writeconcern.Majority()),
	)
	if transactionError != nil {
		return nil, transactionError
	}
	if converted, ok := result.(*QueueFileResult); ok {
		return converted, nil
	} else {
		return nil, errors.New(constants.ERROR_MESSAGES.FailedToConvertData)
	}
}
