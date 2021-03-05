package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/middleware/cors"
)

func main() {
	database.Connect()
	app := fiber.New()
	app.User(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	routes.Setup(app)
	app.Listen(":8000")
}