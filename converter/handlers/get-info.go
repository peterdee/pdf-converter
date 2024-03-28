package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"

	"converter/constants"
	"converter/database"
)

func GetInfo(uid string) (*GetInfoResult, error) {
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

			var queuedEntries int64 = 0
			if queueEntry.Status == constants.QUEUE_STATUSES.Queued {
				count, countError := database.Queue.CountDocuments(
					tctx,
					bson.D{
						{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
						{Key: "updatedAt", Value: bson.D{
							{Key: "$lte", Value: queueEntry.UpdatedAt},
						}},
					},
				)
				if countError != nil {
					return nil, countError
				}
				queuedEntries = count
			}

			jsonData, jsonError := json.Marshal(queueEntry)
			if jsonError != nil {
				return nil, jsonError
			}

			return &GetInfoResult{
				JSON:          string(jsonData),
				QueuePosition: queuedEntries,
			}, nil
		},
		options.Transaction().SetWriteConcern(writeconcern.Majority()),
	)
	if transactionError != nil {
		return nil, transactionError
	}
	if converted, ok := result.(*GetInfoResult); ok {
		return converted, nil
	} else {
		return nil, errors.New(constants.ERROR_MESSAGES.FailedToConvertData)
	}
}
