package constants

type EnvNames struct {
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
}

type QueueStatuses struct {
	InProgress string
	Processed  string
	Queued     string
}
