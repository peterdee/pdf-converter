package download

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v3"

	"api-interface/constants"
	grpc_generated "api-interface/grpc"
)

func DownloadController(context fiber.Ctx) error {
	uid := context.Params("uid", "")
	if uid == "" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_MESSAGES.MissingUID,
		)
	}

	// TODO: hanlde errors properly
	response, downloadError := grpc_generated.DownloadArchive(uid)
	if downloadError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	bytesData, decodeError := hex.DecodeString(response.Bytes)
	if decodeError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	partials := strings.Split(response.Filename, ".")
	filename := strings.Join(partials[:len(partials)-1], ".")
	context.Set(
		"Content-Disposition",
		fmt.Sprintf("inline; filename=%s.zip", filename),
	)

	return context.SendStream(bytes.NewReader(bytesData))
}
