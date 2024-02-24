package utilities

import "github.com/gofiber/fiber/v3"

type ResponseOptions struct {
	Context fiber.Ctx
	Data    interface{}
	Info    string
	Status  int
}
