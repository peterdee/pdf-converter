package grpc_generated

import (
	context "context"
	"log"
	"net"
	"os"
	"strconv"

	grpc "google.golang.org/grpc"

	"converter/constants"
	"converter/handlers"
)

type server struct {
	UnimplementedConverterServer
}

func (s *server) DeleteEntry(
	ctx context.Context,
	in *DeleteEntryRequest,
) (*DeleteEntryResponse, error) {
	result, handlerError := handlers.DeleteEntry(in.Uid)
	if handlerError != nil {
		return nil, handlerError
	}
	return &DeleteEntryResponse{Deleted: result.Deleted}, nil
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
		Count:          getInfoResult.QueuedItems,
		DownloadedAt:   getInfoResult.DownloadedAt,
		Filename:       getInfoResult.Filename,
		Status:         getInfoResult.Status,
		TotalDownloads: int64(getInfoResult.TotalDownloads),
		Uid:            getInfoResult.UID,
	}
	return &response, nil
}

func (s *server) GetQueue(
	ctx context.Context,
	in *GetQueueRequest,
) (*GetQueueResponse, error) {
	getQueueResult, getQueueError := handlers.GetQueue(in.Limit, in.Page)
	if getQueueError != nil {
		return nil, getQueueError
	}
	return &GetQueueResponse{Json: getQueueResult.JSON}, nil
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
		Uid:      queueFileResult.UID,
	}
	return &response, nil
}

func (s *server) TestConnection(
	ctx context.Context,
	in *TestConnectionRequest,
) (*TestConnectionResponse, error) {
	testConnectionResult := handlers.TestConnection(in.Timestamp)
	return &TestConnectionResponse{
		Timestamp: testConnectionResult.Timestamp,
	}, nil
}

func CreateServer(tcpServer net.Listener) {
	maxBodyLimit := constants.MAX_BODY_LIMIT_MB
	maxBodyLimitString := os.Getenv(constants.ENV_NAMES.MaxBodyLimitMB)
	if maxBodyLimitString != "" {
		value, conversionError := strconv.Atoi(maxBodyLimitString)
		if conversionError != nil {
			maxBodyLimit = value
		}
	}
	rpcServer := grpc.NewServer(
		grpc.MaxRecvMsgSize(1024*1024*maxBodyLimit),
		grpc.MaxSendMsgSize(1024*1024*maxBodyLimit),
	)
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
