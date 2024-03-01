package queue

import (
	"github.com/gofiber/fiber/v3"

	"api-interface/constants"
	grpc_client "api-interface/grpc"
	"api-interface/utilities"
)

func GetInfoController(context fiber.Ctx) error {
	uid := context.Params("uid", "")
	if uid == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_MESSAGES.MissingUID,
		)
	}

	info, infoError := grpc_client.GetInfo(uid)
	if infoError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data: fiber.Map{
			"filename":   &info.Filename,
			"queueCount": &info.Count,
			"status":     &info.Status,
			"uid":        &info.Uid,
		},
	})
}
