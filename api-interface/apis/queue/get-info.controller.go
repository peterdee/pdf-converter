package queue

import (
	"encoding/json"
	"strings"

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
		if strings.Contains(infoError.Error(), "INVALID_UID") {
			return fiber.NewError(fiber.StatusBadRequest, constants.RESPONSE_MESSAGES.InvalidUID)
		}
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	var entry QueueEntry
	parsingError := json.Unmarshal([]byte(info.Json), &entry)
	if parsingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data: fiber.Map{
			"createdAt":        entry.CreatedAt,
			"downloadedAt":     entry.DownloadedAt,
			"id":               entry.ID,
			"originalFileName": entry.OriginalFileName,
			"queuePosition":    info.QueuePosition,
			"status":           entry.Status,
			"totalDownloads":   entry.TotalDownloads,
			"uid":              entry.UID,
			"updatedAt":        entry.UpdatedAt,
		},
	})
}
