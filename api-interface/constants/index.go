package constants

var ERROR_MESSAGES = ErrorMessages{
	CouldNotConnectToRPC:       "Could not connect to RPC",
	CouldNotLoadEnvFile:        "Could not load .env file",
	CouldNotLoadRPCCredentials: "Could not load RPC credentials",
}

var RESPONSE_MESSAGES = ResponseMessages{
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidFile:         "INVALID_FILE",
	MissingFile:         "MISSING_FILE",
	Ok:                  "OK",
}

var PORT = "8000"
