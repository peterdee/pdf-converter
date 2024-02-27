package grpc_client

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

var Client SplitterClient
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

	Client = NewSplitterClient(Connection)
}

func SplitFile(bytes []byte, filename string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, splitDataError := Client.SplitData(
		ctx,
		&SplitDataRequest{Bytes: hex.EncodeToString(bytes), Filename: filename},
	)
	if splitDataError != nil {
		return "", splitDataError
	}
	return response.Id, nil
}
