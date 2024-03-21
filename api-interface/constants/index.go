package constants

var ENV_NAMES = EnvNames{
	GRPCAddress:      "GRPC_ADDRESS",
	GRPCMaxReceiveMB: "GRPC_MAX_RECEIVE_MB",
	MaxBodyLimitMB:   "MAX_BODY_LIMIT_MB",
	Port:             "PORT",
}

var ERROR_MESSAGES = ErrorMessages{
	CouldNotConnectToRPC:       "Could not connect to RPC",
	CouldNotLoadEnvFile:        "Could not load .env file",
	CouldNotLoadRPCCredentials: "Could not load RPC credentials",
}

const GRPC_MAX_RECEIVE_MB int = 500

const MAX_BODY_LIMIT_MB int = 150

var RESPONSE_MESSAGES = ResponseMessages{
	InternalServerError: "INTERNAL_SERVER_ERROR",
	InvalidFile:         "INVALID_FILE",
	InvalidUID:          "INVALID_UID",
	MissingFile:         "MISSING_FILE",
	MissingUID:          "MISSING_UID",
	Ok:                  "OK",
}

var PORT = "8000"
