package index

import (
	"github.com/gofiber/fiber/v3"

	"api-interface/utilities"
)

func IndexController(context fiber.Ctx) error {
	return utilities.Response(utilities.ResponseOptions{
		Context: context,
	})
}
