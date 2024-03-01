package queue

import "github.com/gofiber/fiber/v3"

func RegisterControllers(app *fiber.App) {
	group := app.Group("/api/queue")

	group.Post("/", QueueFileController)
	group.Get("/:uid", GetInfoController)
}
