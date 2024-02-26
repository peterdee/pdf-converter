package grpc_client

import (
	context "context"
	"fmt"
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
	defer connection.Close()

	Connection = connection

	Client = NewSplitterClient(Connection)
}

func SplitFile(bytes []byte, filename string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := Client.SplitData(ctx, &SplitDataRequest{Bytes: string(bytes), Filename: filename})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	fmt.Println(r)
}
