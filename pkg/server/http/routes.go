package http

import (
	"github.com/gofiber/fiber"
)

func attachExternalRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World!")
	})
}

func attachInternalRoutes(app *fiber.App) {

}
