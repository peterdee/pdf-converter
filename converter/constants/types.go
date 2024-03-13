package constants

type EnvNames struct {
	ArchiveStoredDays     string
	MaxBodyLimit          string
	MongoConnectionString string
	MongoDatabaseName     string
	Port                  string
}

type ErrorMessages struct {
	FailedToServeRPCServer            string
	CouldNotConnectToMongo            string
	CouldNotConnectToRPC              string
	CouldNotLoadEnvFile               string
	CouldNotLoadMongoConnectionString string
	CouldNotLoadRPCCredentials        string
	MongoTransactionError             string
	TargetDirectoryIsEmpty            string
}

type QueueStatuses struct {
	InProgress string
	Processed  string
	Queued     string
}

type ResponseErrors struct {
	ArchiveAlreadyDeleted string
	InvalidUID            string
	NotProcessed          string
}
