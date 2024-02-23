package main

import (
	"api-interface/constants"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/favicon"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/joho/godotenv"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal(constants.ERROR_MESSAGES.CouldNotLoadEnvFile)
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: nil,
	})

	app.Use(favicon.New(favicon.Config{
		File: "./assets/favicon.ico",
	}))
	app.Use(logger.New())

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("response")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = constants.PORT
	}
	if launchError := app.Listen(fmt.Sprintf(":%s", port)); launchError != nil {
		log.Fatal(launchError)
	}
}
