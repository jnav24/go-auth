package routes

import (
	"../controllers"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *Fiber.App) {
	app.Get("/", controllers.Index)
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Get("/api/logout", controllers.Logout)
}