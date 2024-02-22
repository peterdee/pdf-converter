package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {
	if envError := godotenv.Load(); envError != nil {
		log.Fatal("Could not load .env file")
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: nil,
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("response")
	})

	app.Listen(fmt.Sprintf(":%s", os.Getenv("PORT")))
}
