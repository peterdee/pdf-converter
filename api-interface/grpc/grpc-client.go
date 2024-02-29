package grpc_generated

import (
	context "context"
	"encoding/hex"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"api-interface/constants"
)

var Client ConverterClient
var Connection *grpc.ClientConn

func CreateRPCConnection() {
	address := os.Getenv("GRPC_ADDRESS")
	if address == "" {
		log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadRPCCredentials)
	}
	connection, connectionError := grpc.Dial(
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if connectionError != nil {
		log.Fatalf(
			"%s: %v",
			constants.ERROR_MESSAGES.CouldNotConnectToRPC,
			connectionError,
		)
	}

	Connection = connection

	Client = NewConverterClient(Connection)
}

func QueueFile(bytes []byte, filename string) (*QueueFileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, splitDataError := Client.QueueFile(
		ctx,
		&QueueFileRequest{Bytes: hex.EncodeToString(bytes), Filename: filename},
	)
	if splitDataError != nil {
		return nil, splitDataError
	}
	return response, nil
}
