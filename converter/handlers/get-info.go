package handlers

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"converter/constants"
	"converter/database"
)

func GetInfo(uid string) (*GetInfoResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var queueEntry database.QueueEntry
	queryError := database.Queue.FindOne(
		ctx,
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
			ctx,
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

	response := &GetInfoResult{
		Filename:    queueEntry.OriginalFileName,
		QueuedItems: queuedEntries,
		Status:      queueEntry.Status,
		UID:         uid,
	}
	return response, nil
}
