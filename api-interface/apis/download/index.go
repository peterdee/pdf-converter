package download

import "github.com/gofiber/fiber/v3"

func RegisterControllers(app *fiber.App) {
	group := app.Group("/api/download")

	group.Get("/:uid", DownloadController)
}
