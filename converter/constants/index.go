package constants

const DATABASE_NAME = "queuedb"

var ENV_NAMES = EnvNames{
	MongoConnectionString: "MONGO_CONNECTION_STRING",
	MongoDatabaseName:     "MONGO_DATABASE_NAME",
	Port:                  "PORT",
}

var ERROR_MESSAGES = ErrorMessages{
	FailedToServeRPCServer:            "Failed to serve gRPC server",
	CouldNotConnectToMongo:            "Could not connect to MongoDB",
	CouldNotConnectToRPC:              "Could not connect to RPC",
	CouldNotLoadEnvFile:               "Could not load .env file",
	CouldNotLoadMongoConnectionString: "Could not load MongoDB connection string",
	CouldNotLoadRPCCredentials:        "Could not load RPC credentials",
	MongoTransactionError:             "Mongo transaction error",
	TargetDirectoryIsEmpty:            "target directory is empty",
}

const PORT = "50051"

var QUEUE_STATUSES = QueueStatuses{
	InProgress: "in-progress",
	Processed:  "processed",
	Queued:     "queued",
}
