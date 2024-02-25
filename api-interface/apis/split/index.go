package split

import "github.com/gofiber/fiber/v3"

func RegisterControllers(app *fiber.App) {
	group := app.Group("/api/split")

	group.Post("/", SplitController)
}
