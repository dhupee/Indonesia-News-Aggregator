package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// Load the .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load .env")
	}

	// Create a new Fiber app
	app := fiber.New()

	// Get the port from the environment variables
	port := os.Getenv("PORT")

	// Define the route handler for the root path
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	// search news on Kompas
	app.Get("/kompas/search/:keyword?", func(c *fiber.Ctx) error {
		return c.SendString("You search for " + c.Params("keyword"))
	})

	app.Get("/kompas/categories/:category?", func(c *fiber.Ctx) error {
		return c.SendString("You search for " + c.Params("category") + " category")
	})


	// Start the app on the specified port
	app.Listen(":" + port)
}
