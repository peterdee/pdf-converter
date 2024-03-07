package constants

type EnvNames struct {
	GRPCAddress  string
	MaxBodyLimit string
	Port         string
}

type ErrorMessages struct {
	CouldNotConnectToRPC       string
	CouldNotLoadEnvFile        string
	CouldNotLoadRPCCredentials string
}

type ResponseMessages struct {
	InternalServerError string
	InvalidFile         string
	InvalidUID          string
	MissingFile         string
	MissingUID          string
	Ok                  string
}
