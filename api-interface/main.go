package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
	"github.com/julyskies/gohelpers"

	"api-interface/apis/download"
	"api-interface/apis/index"
	"api-interface/apis/queue"
	"api-interface/constants"
	grpc_client "api-interface/grpc"
	"api-interface/utilities"
)

func main() {
	envSource := os.Getenv(constants.ENV_NAMES.EnvSource)
	if envSource != "env" {
		if envError := godotenv.Load(); envError != nil {
			log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadEnvFile)
		}
	}

	grpc_client.CreateRPCConnection()
	for i := 0; i <= 5; i += 1 {
		timestamp := gohelpers.MakeTimestamp()
		_, testError := grpc_client.TestConnection(timestamp)
		if testError != nil {
			if i == 5 {
				log.Fatal(constants.ERROR_MESSAGES.CouldNotConnectToRPC)
			}
			log.Printf("Could not connect to RPC, retrying in %d sec", i+1)
			time.Sleep(time.Second * time.Duration(i+1))
		}
	}

	maxBodyLimit := constants.MAX_BODY_LIMIT_MB
	maxBodyLimitString := os.Getenv(constants.ENV_NAMES.MaxBodyLimitMB)
	if maxBodyLimitString != "" {
		value, conversionError := strconv.Atoi(maxBodyLimitString)
		if conversionError == nil {
			maxBodyLimit = value
		}
	}
	app := fiber.New(fiber.Config{
		BodyLimit:    maxBodyLimit * 1024 * 1024,
		ErrorHandler: utilities.ErrorHandler,
	})

	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(logger.New())

	download.RegisterControllers(app)
	index.RegisterControllers(app)
	queue.RegisterControllers(app)

	port := os.Getenv(constants.ENV_NAMES.Port)
	if port == "" {
		port = constants.PORT
	}
	if launchError := app.Listen(fmt.Sprintf(":%s", port)); launchError != nil {
		log.Fatal(launchError)
	}
}
