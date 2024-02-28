package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"

	"splitter/constants"
	"splitter/database"
	grpc_client "splitter/grpc"
	"splitter/utilities"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadEnvFile)
	}

	database.CreateConnection()
	utilities.LaunchCRON()

	if mkdirError := os.Mkdir("./processing", 0777); mkdirError != nil {
		if mkdirError.Error() != "mkdir ./processing: file exists" {
			log.Fatal(mkdirError)
		}
	}

	port := os.Getenv(constants.ENV_NAMES.Port)
	if port == "" {
		port = constants.PORT
	}
	tcpServer, tcpError := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if tcpError != nil {
		log.Fatal(tcpError)
	}
	grpc_client.CreateServer(tcpServer)
}
