package constants

const ARCHIVE_STORED_DAYS int = 14

const DATABASE_NAME = "queuedb"

var ENV_NAMES = EnvNames{
	ArchiveStoredDays:     "ARCHIVE_STORED_DAYS",
	MaxBodyLimit:          "MAX_BODY_LIMIT",
	MongoConnectionString: "MONGO_CONNECTION_STRING",
	MongoDatabaseName:     "MONGO_DATABASE_NAME",
	Port:                  "PORT",
}

var ERROR_MESSAGES = ErrorMessages{
	CouldNotConnectToMongo:            "Could not connect to MongoDB",
	CouldNotConnectToRPC:              "Could not connect to RPC",
	CouldNotLoadEnvFile:               "Could not load .env file",
	CouldNotLoadMongoConnectionString: "Could not load MongoDB connection string",
	CouldNotLoadRPCCredentials:        "Could not load RPC credentials",
	FailedToServeRPCServer:            "Failed to serve gRPC server",
	MongoTransactionError:             "Mongo transaction error",
	TargetDirectoryIsEmpty:            "target directory is empty",
}

const MAX_BODY_LIMIT int = 300

const PORT = "50051"

var QUEUE_STATUSES = QueueStatuses{
	InProgress: "in-progress",
	Processed:  "processed",
	Queued:     "queued",
}

var RESPONSE_ERRORS = ResponseErrors{
	ArchiveAlreadyDeleted: "ARCHIVE_ALREADY_DELETED",
	InvalidUID:            "INVALID_UID",
	NotProcessed:          "NOT_PROCESSED",
}
