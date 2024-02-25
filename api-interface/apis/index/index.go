package index

import "github.com/gofiber/fiber/v3"

func RegisterControllers(app *fiber.App) {
	group := app.Group("/")

	group.Get("/", IndexController)
	group.Get("/api", IndexController)
}
