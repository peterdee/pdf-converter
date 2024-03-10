package handlers

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"converter/constants"
	"converter/database"
)

func DownloadArchive(uid string) (*DownloadArchiveResult, error) {
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

	// TODO: should delete files on some condition
	// defer func() {
	// 	os.Remove(filePath)
	// }()

	response := &DownloadArchiveResult{
		Bytes:       hex.EncodeToString(bytes),
		Filename:    queueEntry.OriginalFileName,
		ProcessedAt: queueEntry.UpdatedAt,
		UID:         queueEntry.UID,
	}
	return response, nil
}
