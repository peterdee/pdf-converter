package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"

	"api-interface/apis/download"
	"api-interface/apis/index"
	"api-interface/apis/queue"
	"api-interface/constants"
	grpc_client "api-interface/grpc"
	"api-interface/utilities"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadEnvFile)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: utilities.ErrorHandler,
	})

	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(logger.New())

	grpc_client.CreateRPCConnection()

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
