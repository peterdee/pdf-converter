package database

import "go.mongodb.org/mongo-driver/bson/primitive"

type QueueEntry struct {
	CreatedAt        int64              `json:"createdAt"`
	DownloadedAt     int64              `json:"downloadedAt"`
	ID               primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	OriginalFileName string             `json:"originalFileName"`
	Status           string             `json:"status"`
	UID              string             `json:"uid"`
	UpdatedAt        int64              `json:"updatedAt"`
}
