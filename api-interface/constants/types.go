package constants

type EnvNames struct {
	EnvSource        string
	GRPCAddress      string
	GRPCMaxReceiveMB string
	MaxBodyLimitMB   string
	Port             string
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
