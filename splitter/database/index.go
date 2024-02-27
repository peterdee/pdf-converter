package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

var Database *mongo.Database

var Queue *mongo.Collection

func CreateConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionString := os.Getenv("MONGO_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatalf("Could not load MongoDB connection string")
	}
	client, connectionError := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if connectionError != nil {
		log.Fatalf("Could not connect to MongoDB: %v", connectionError)
	}

	Client = client
	Database = Client.Database("splitqueue")
	Queue = Database.Collection("queue")
}
