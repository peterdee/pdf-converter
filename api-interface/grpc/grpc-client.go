package grpc_generated

import (
	context "context"
	"encoding/hex"
	"fmt"
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
	address := os.Getenv(constants.ENV_NAMES.GRPCAddress)
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

func DownloadArchive(uid string) (*DownloadArchiveResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, responseError := Client.DownloadArchive(
		ctx,
		&DownloadArchiveRequest{Uid: uid},
	)
	if responseError != nil {
		return nil, responseError
	}
	return response, nil
}

func GetInfo(uid string) (*GetInfoResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, responseError := Client.GetInfo(
		ctx,
		&GetInfoRequest{Uid: uid},
	)
	if responseError != nil {
		return nil, responseError
	}
	return response, nil
}

func QueueFile(bytes []byte, filename string) (*QueueFileResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	response, responseError := Client.QueueFile(
		ctx,
		&QueueFileRequest{Bytes: hex.EncodeToString(bytes), Filename: filename},
	)
	if responseError != nil {
		fmt.Println(responseError)
		return nil, responseError
	}
	return response, nil
}
