package constants

type ErrorMessages struct {
	CouldNotConnectToRPC       string
	CouldNotLoadEnvFile        string
	CouldNotLoadRPCCredentials string
}

type ResponseMessages struct {
	InternalServerError string
	InvalidFile         string
	MissingFile         string
	Ok                  string
}
