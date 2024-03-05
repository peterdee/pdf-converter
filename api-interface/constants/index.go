package constants

var ENV_NAMES = EnvNames{
	GRPCAddress: "GRPC_ADDRESS",
	Port:        "PORT",
}

var ERROR_MESSAGES = ErrorMessages{
	CouldNotConnectToRPC:       "Could not connect to RPC",
	CouldNotLoadEnvFile:        "Could not load .env file",
	CouldNotLoadRPCCredentials: "Could not load RPC credentials",
}

const MAX_BODY_LIMIT int = 150

var RESPONSE_MESSAGES = ResponseMessages{
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidFile:         "INVALID_FILE",
	MissingFile:         "MISSING_FILE",
	MissingUID:          "MISSING_UID",
	Ok:                  "OK",
}

var PORT = "8000"
