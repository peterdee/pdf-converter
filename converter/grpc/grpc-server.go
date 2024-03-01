package grpc_generated

import (
	context "context"
	"log"
	"net"

	grpc "google.golang.org/grpc"

	"converter/constants"
	"converter/handlers"
)

type server struct {
	UnimplementedConverterServer
}

func (s *server) DownloadArchive(
	ctx context.Context,
	in *DownloadArchiveRequest,
) (*DownloadArchiveResponse, error) {
	result, downloadError := handlers.DownloadArchive(in.Uid)
	if downloadError != nil {
		return nil, downloadError
	}
	response := DownloadArchiveResponse{
		Bytes:     result.Bytes,
		Filename:  result.Filename,
		Processed: result.ProcessedAt,
		Uid:       result.UID,
	}
	return &response, nil
}

func (s *server) GetInfo(
	ctx context.Context,
	in *GetInfoRequest,
) (*GetInfoResponse, error) {
	getInfoResult, getInfoError := handlers.GetInfo(in.Uid)
	if getInfoError != nil {
		return nil, getInfoError
	}
	response := GetInfoResponse{
		Count:    getInfoResult.QueuedItems,
		Filename: getInfoResult.Filename,
		Status:   getInfoResult.Status,
		Uid:      getInfoResult.UID,
	}
	return &response, nil
}

func (s *server) QueueFile(
	ctx context.Context,
	in *QueueFileRequest,
) (*QueueFileResponse, error) {
	queueFileResult, queueError := handlers.QueueFile(in.Bytes, in.Filename)
	if queueError != nil {
		return nil, queueError
	}
	response := QueueFileResponse{
		Count:    queueFileResult.QueuedItems,
		Filename: queueFileResult.Filename,
		Status:   queueFileResult.Status,
		Uid:      queueFileResult.QueueId,
	}
	return &response, nil
}

func CreateServer(tcpServer net.Listener) {
	rpcServer := grpc.NewServer()
	RegisterConverterServer(rpcServer, &server{})
	log.Printf("gRPC server is listening at %v", tcpServer.Addr())
	if serveError := rpcServer.Serve(tcpServer); serveError != nil {
		log.Fatalf(
			"%s: %v",
			constants.ERROR_MESSAGES.FailedToServeRPCServer,
			serveError,
		)
	}
}
