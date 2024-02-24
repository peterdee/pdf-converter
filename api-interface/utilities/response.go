package utilities

import (
	"github.com/gofiber/fiber/v3"
	"github.com/julyskies/gohelpers"

	"api-interface/constants"
)

func Response(options ResponseOptions) error {
	info := options.Info
	if info == "" {
		info = constants.RESPONSE_MESSAGES.Ok
	}
	status := options.Status
	if status == 0 {
		status = fiber.StatusOK
	}
	response := fiber.Map{
		"datetime": gohelpers.MakeTimestamp(),
		"info":     info,
		"request":  options.Context.OriginalURL() + " [" + options.Context.Method() + "]",
		"status":   status,
	}
	if options.Data != nil {
		response["data"] = options.Data
	}
	return options.Context.Status(status).JSON(response)
}
