package handlers

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/julyskies/gohelpers"
	"go.mongodb.org/mongo-driver/bson"

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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	now := gohelpers.MakeTimestamp()
	_, insertionError := database.Queue.InsertOne(ctx, bson.D{
		{Key: "createdAt", Value: now},
		{Key: "downloadedAt", Value: 0},
		{Key: "originalFileName", Value: filename},
		{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
		{Key: "uid", Value: hash},
		{Key: "updatedAt", Value: now},
	})
	if insertionError != nil {
		return nil, insertionError
	}

	count, countError := database.Queue.CountDocuments(
		ctx,
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
}
