package database

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"converter/constants"
)

const QueueCollection string = "Queue"

var Client *mongo.Client

var Database *mongo.Database

var Queue *mongo.Collection

func CreateConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	connectionString := os.Getenv(constants.ENV_NAMES.MongoConnectionString)
	if connectionString == "" {
		log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadMongoConnectionString)
	}
	databaseName := os.Getenv(constants.ENV_NAMES.MongoDatabaseName)
	if databaseName == "" {
		databaseName = constants.DATABASE_NAME
	}
	client, connectionError := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if connectionError != nil {
		log.Fatalf(
			"%s: %v",
			constants.ERROR_MESSAGES.CouldNotConnectToMongo,
			connectionError,
		)
	}

	log.Println("MongoDB connected")

	Client = client
	Database = Client.Database(databaseName)
	Queue = Database.Collection(QueueCollection)
}
