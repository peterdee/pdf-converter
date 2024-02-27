package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	"splitter/constants"
	grpc_client "splitter/grpc"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadEnvFile)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = constants.PORT
	}
	tcpServer, tcpError := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if tcpError != nil {
		log.Fatal(tcpError)
	}

	grpc_client.CreateServer(tcpServer)
}
