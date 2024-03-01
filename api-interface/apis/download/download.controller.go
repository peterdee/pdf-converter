package download

import (
	"github.com/gofiber/fiber/v3"

	"api-interface/constants"
	grpc_generated "api-interface/grpc"
	"api-interface/utilities"
)

func DownloadController(context fiber.Ctx) error {
	uid := context.Params("uid", "")
	if uid == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_MESSAGES.MissingUID,
		)
	}

	response, downloadError := grpc_generated.DownloadArchive(uid)
	if downloadError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	// TODO: should download archive

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data: fiber.Map{
			"bytes":       &response.Bytes,
			"filename":    &response.Filename,
			"processedAt": &response.Processed,
			"uid":         &response.Uid,
		},
	})
}
