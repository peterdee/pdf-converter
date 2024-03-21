package handlers

import (
	"context"
	"converter/constants"
	"converter/database"
	"encoding/json"
	"errors"
	"math"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

const LIMIT int64 = 20
const MAX_LIMIT int64 = 100
const PAGE int64 = 1

func GetQueue(limit, page int64) (*GetQueueResult, error) {
	if limit == 0 {
		limit = LIMIT
	}
	if limit > MAX_LIMIT {
		limit = MAX_LIMIT
	}
	if page == 0 {
		page = PAGE
	}

	offset := (page - 1) * limit

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
			cursor, cursorError := database.Queue.Find(
				tctx,
				bson.D{},
				&options.FindOptions{
					Limit: &limit,
					Skip:  &offset,
				},
			)
			if cursorError != nil && cursorError != mongo.ErrNoDocuments {
				return nil, cursorError
			}

			var documents []database.QueueEntry
			decodeError := cursor.All(tctx, &documents)
			if decodeError != nil {
				return nil, decodeError
			}

			totalCount, countError := database.Queue.CountDocuments(
				tctx,
				bson.D{},
			)
			if countError != nil {
				return nil, countError
			}

			paginatedResponse := &GetQueueJSON{
				Items:      documents,
				Limit:      int(limit),
				Page:       int(page),
				TotalItems: int(totalCount),
				TotalPages: int(math.Ceil(float64(totalCount) / float64(limit))),
			}
			jsonResponse, jsonError := json.Marshal(paginatedResponse)
			if jsonError != nil {
				return nil, jsonError
			}
			return &GetQueueResult{JSON: string(jsonResponse)}, nil
		},
		options.Transaction().SetWriteConcern(writeconcern.Majority()),
	)
	if transactionError != nil {
		return nil, transactionError
	}
	if converted, ok := result.(*GetQueueResult); ok {
		return converted, nil
	} else {
		return nil, errors.New(constants.ERROR_MESSAGES.FailedToConvertData)
	}
}
