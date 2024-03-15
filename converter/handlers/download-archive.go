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
)

func DownloadArchive(uid string) (*DownloadArchiveResult, error) {
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
			var queueEntry database.QueueEntry
			queryError := database.Queue.FindOne(
				tctx,
				bson.D{{Key: "uid", Value: uid}},
			).Decode(&queueEntry)
			if queryError != nil && queryError != mongo.ErrNoDocuments {
				return nil, queryError
			}
			if queueEntry.UID == "" {
				return nil, errors.New(constants.RESPONSE_ERRORS.InvalidUID)
			}
			if queueEntry.Status != constants.QUEUE_STATUSES.Processed {
				return nil, errors.New(constants.RESPONSE_ERRORS.NotProcessed)
			}

			filePath := fmt.Sprintf("./processing/%s/%s.zip", queueEntry.UID, queueEntry.UID)
			_, fileError := os.Stat(filePath)
			if fileError != nil {
				return nil, errors.New(constants.RESPONSE_ERRORS.ArchiveAlreadyDeleted)
			}

			bytes, readError := os.ReadFile(filePath)
			if readError != nil {
				return nil, errors.New(constants.RESPONSE_ERRORS.ArchiveAlreadyDeleted)
			}

			_, queryError = database.Queue.UpdateOne(
				tctx,
				bson.D{{Key: "uid", Value: queueEntry.UID}},
				bson.D{
					{Key: "$set", Value: bson.D{
						{Key: "downloadedAt", Value: gohelpers.MakeTimestamp()},
					}},
				},
			)
			if queryError != nil && queryError != mongo.ErrNoDocuments {
				return nil, queryError
			}

			response := &DownloadArchiveResult{
				Bytes:       hex.EncodeToString(bytes),
				Filename:    queueEntry.OriginalFileName,
				ProcessedAt: queueEntry.UpdatedAt,
				UID:         queueEntry.UID,
			}
			return response, nil
		},
		options.Transaction().SetWriteConcern(writeconcern.Majority()),
	)
	if transactionError != nil {
		return nil, transactionError
	}
	if converted, ok := result.(*DownloadArchiveResult); ok {
		return converted, nil
	} else {
		return nil, errors.New(constants.ERROR_MESSAGES.FailedToConvertData)
	}
}
