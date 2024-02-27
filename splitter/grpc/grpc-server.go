package grpc_client

import (
	context "context"
	"log"
	"net"
	"splitter/handlers"

	grpc "google.golang.org/grpc"
)

type server struct {
	UnimplementedSplitterServer
}

func (s *server) SplitData(ctx context.Context, in *SplitDataRequest) (*SplitDataResponse, error) {
	queueId, queueError := handlers.SplitPDF(in.Bytes, in.Filename)
	if queueError != nil {
		return nil, queueError
	}
	return &SplitDataResponse{Id: queueId}, nil
}

func CreateServer(tcpServer net.Listener) {
	rpcServer := grpc.NewServer()
	RegisterSplitterServer(rpcServer, &server{})
	log.Printf("gRPC server is listening at %v", tcpServer.Addr())
	if serveError := rpcServer.Serve(tcpServer); serveError != nil {
		log.Fatalf("Failed to serve gRPC server: %v", serveError)
	}
}
