package main

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func main() {
	// change username, password, and db credentials
	_, err := gorm.Open(mysql.Open("username:password@/database"), &gorm.Config({}}))

	if err != nil {
		panic("could not connect to the database")
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Listen(":8000")
}