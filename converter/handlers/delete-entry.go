package handlers

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"converter/database"
)

func DeleteEntry(uid string) (*DeleteEntryResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, deleteError := database.Queue.DeleteOne(ctx, bson.D{{Key: "uid", Value: uid}})
	if deleteError != nil {
		return nil, deleteError
	}
	if result.DeletedCount > 0 {
		os.RemoveAll(fmt.Sprintf("./processing/%s", uid))
	}

	return &DeleteEntryResult{Deleted: result.DeletedCount > 0}, nil
}
