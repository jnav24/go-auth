package routes

import "github.com/gofiber/fiber/v2"

func Setup(app *Fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}