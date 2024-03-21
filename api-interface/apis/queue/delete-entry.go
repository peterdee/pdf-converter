package queue

import (
	"github.com/gofiber/fiber/v3"

	"api-interface/constants"
	grpc_client "api-interface/grpc"
	"api-interface/utilities"
)

func DeleteEntryController(context fiber.Ctx) error {
	uid := context.Params("uid", "")
	if uid == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_MESSAGES.MissingUID,
		)
	}

	result, deleteError := grpc_client.DeleteEntry(uid)
	if deleteError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	if !result.Deleted {
		return fiber.NewError(fiber.StatusBadRequest, constants.RESPONSE_MESSAGES.InvalidUID)
	}

	return utilities.Response(utilities.ResponseOptions{Context: context})
}
