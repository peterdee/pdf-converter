package grpc_generated

import (
	context "context"
	"encoding/hex"
	"log"
	"os"
	"strconv"
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

func DeleteEntry(uid string) (*DeleteEntryResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	response, responseError := Client.DeleteEntry(
		ctx,
		&DeleteEntryRequest{Uid: uid},
	)
	if responseError != nil {
		return nil, responseError
	}
	return response, nil
}

func DownloadArchive(uid string) (*DownloadArchiveResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	maxReceiveMB := constants.GRPC_MAX_RECEIVE_MB
	maxReceiveMBString := os.Getenv(constants.ENV_NAMES.GRPCMaxReceiveMB)
	if maxReceiveMBString != "" {
		value, convertError := strconv.Atoi(maxReceiveMBString)
		if convertError == nil {
			maxReceiveMB = value
		}
	}

	response, responseError := Client.DownloadArchive(
		ctx,
		&DownloadArchiveRequest{Uid: uid},
		grpc.MaxCallRecvMsgSize(maxReceiveMB*1024*1024),
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

func GetQueue(limit, page int64) (*GetQueueResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, responseError := Client.GetQueue(
		ctx,
		&GetQueueRequest{Limit: limit, Page: page},
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
		return nil, responseError
	}
	return response, nil
}

func TestConnection(timestamp int64) (*TestConnectionResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	response, responseError := Client.TestConnection(
		ctx,
		&TestConnectionRequest{Timestamp: timestamp},
	)
	if responseError != nil {
		return nil, responseError
	}
	return response, nil
}
