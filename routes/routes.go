package routes

import (
	"../controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *Fiber.App) {
	app.Get("/", controllers.Index)
	app.Post("/api/register", controllers.Register)
}