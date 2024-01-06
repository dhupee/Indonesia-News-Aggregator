package main

import (
	// "fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// create new fiber instance
	app := fiber.New()

	// home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start server
	app.Listen(":3000")
}