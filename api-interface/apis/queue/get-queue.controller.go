package queue

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v3"

	grpc_client "api-interface/grpc"
	"api-interface/utilities"
)

const LIMIT int = 20
const MAX_LIMIT int = 100
const PAGE int = 1

func GetQueueController(context fiber.Ctx) error {
	stringLimit := context.Query("limit", fmt.Sprintf("%d", LIMIT))
	stringPage := context.Query("page", fmt.Sprintf("%d", PAGE))
	limit, convertError := strconv.Atoi(stringLimit)
	if convertError != nil {
		limit = LIMIT
	}
	if limit > MAX_LIMIT {
		limit = MAX_LIMIT
	}
	page, convertError := strconv.Atoi(stringPage)
	if convertError != nil {
		page = PAGE
	}

	result, resultError := grpc_client.GetQueue(int64(limit), int64(page))
	if resultError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	var parsedData GetQueueJSON
	parsingError := json.Unmarshal([]byte(result.Json), &parsedData)
	if parsingError != nil {
		return fiber.NewError(fiber.StatusInternalServerError)
	}

	return utilities.Response(utilities.ResponseOptions{
		Context: context,
		Data: fiber.Map{
			"items":      parsedData.Items,
			"limit":      parsedData.Limit,
			"page":       parsedData.Page,
			"totalItems": parsedData.TotalItems,
			"totalPages": parsedData.TotalPages,
		},
	})
}
