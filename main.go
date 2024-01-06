package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Failed to load .env")
	}

	// create new fiber instance
	app := fiber.New()

	// env variables
	PORT := os.Getenv("PORT")

	// home route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// start server
	app.Listen(":" + PORT)
}
