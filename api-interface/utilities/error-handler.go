package utilities

import (
	"api-interface/constants"

	"github.com/gofiber/fiber/v3"
)

func ErrorHandler(context fiber.Ctx, err error) error {
	info := constants.RESPONSE_MESSAGES.InternalServerError
	status := fiber.StatusInternalServerError

	if e, ok := err.(*fiber.Error); ok {
		info = e.Message
		status = e.Code
	}

	return Response(ResponseOptions{
		Context: context,
		Info:    info,
		Status:  status,
	})
}
