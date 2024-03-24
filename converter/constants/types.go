package constants

type EnvNames struct {
	ArchiveStoredDays     string
	EnvSource             string
	MaxBodyLimitMB        string
	MongoConnectionString string
	MongoDatabaseName     string
	Port                  string
}

type ErrorMessages struct {
	CouldNotConnectToMongo            string
	CouldNotConnectToRPC              string
	CouldNotLoadEnvFile               string
	CouldNotLoadMongoConnectionString string
	CouldNotLoadRPCCredentials        string
	FailedToConvertData               string
	FailedToServeRPCServer            string
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
