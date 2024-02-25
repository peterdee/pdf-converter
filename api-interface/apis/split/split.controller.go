package split

import (
	"fmt"

	"github.com/gofiber/fiber/v3"

	"api-interface/utilities"
)

func SplitController(context fiber.Ctx) error {
	file, fileError := context.FormFile("pdf")
	if fileError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}
	fmt.Println(file)

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
	})
}
