package queue

import "github.com/gofiber/fiber/v3"

func RegisterControllers(app *fiber.App) {
	group := app.Group("/api/queue")

	group.Delete("/:uid", DeleteEntryController)
	group.Get("/", GetQueueController)
	group.Get("/:uid", GetInfoController)
	group.Post("/", QueueFileController)
}
