package split

import (
	"io"
	"strings"

	"github.com/gofiber/fiber/v3"

	"api-interface/constants"
	grpc_client "api-interface/grpc"
	"api-interface/utilities"
)

func SplitController(context fiber.Ctx) error {
	file, fileError := context.FormFile("pdf")
	if fileError != nil {
		if fileError.Error() == "there is no uploaded file associated with the given key" &&
			file == nil {
			return fiber.NewError(
				fiber.StatusBadRequest,
				constants.RESPONSE_MESSAGES.MissingFile,
			)
		}
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	partials := strings.Split(file.Filename, ".")
	extension := partials[len(partials)-1]
	if extension != "pdf" {
		return fiber.NewError(
			fiber.StatusBadRequest,
			constants.RESPONSE_MESSAGES.InvalidFile,
		)
	}

	fileContent, openingError := file.Open()
	if openingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	bytes, bytesError := io.ReadAll(fileContent)
	if bytesError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	queueId, queueError := grpc_client.SplitFile(bytes, file.Filename)
	if queueError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data: fiber.Map{
			"queueId": queueId,
		},
	})
}
