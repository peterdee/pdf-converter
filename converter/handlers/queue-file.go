package handlers

import (
	"context"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/julyskies/gohelpers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"converter/constants"
	"converter/database"
	"converter/utilities"
)

func QueueFile(encoded string, filename string) (QueueFileResult, error) {
	hash := utilities.CreateMD5Hash(fmt.Sprintf(
		"%s:%s:%d",
		encoded,
		filename,
		gohelpers.MakeTimestamp(),
	))
	dirPath := fmt.Sprintf("./processing/%s", hash)
	mkdirError := os.Mkdir(dirPath, 0777)
	if mkdirError != nil {
		return QueueFileResult{}, mkdirError
	}

	bytes, decodeError := hex.DecodeString(encoded)
	if decodeError != nil {
		return QueueFileResult{}, decodeError
	}
	os.WriteFile(fmt.Sprintf("%s/%s.pdf", dirPath, hash), bytes, 0777)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	count, countError := database.Queue.EstimatedDocumentCount(ctx)
	if countError != nil {
		return QueueFileResult{}, countError
	}

	now := gohelpers.MakeTimestamp()
	result, insertionError := database.Queue.InsertOne(ctx, bson.D{
		{Key: "createdAt", Value: now},
		{Key: "originalFileName", Value: filename},
		{Key: "status", Value: constants.QUEUE_STATUSES.Queued},
		{Key: "uid", Value: hash},
		{Key: "updatedAt", Value: now},
	})
	if insertionError != nil {
		return QueueFileResult{}, insertionError
	}

	var queueResult = QueueFileResult{
		Filename:    filename,
		QueueId:     result.InsertedID.(primitive.ObjectID).Hex(),
		QueuedItems: count,
		Status:      constants.QUEUE_STATUSES.Queued,
	}
	return queueResult, nil
}
