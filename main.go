package main

import (
	"log"
	"os"

	kompas "github.com/dhupee/Indonesia-News-Aggregator/kompas"

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
		// return welcome.txt

		return c.SendFile("./assets/welcome.txt")
	})

	// search news on Kompas
	app.Get("/kompas/search/:keyword?", func(c *fiber.Ctx) error {
		if c.Params("keyword") == "" {
			return c.SendString("Please specify keyword")
		}

		result := kompas.Search(c.Params("keyword"))
		return c.SendString("You search for " + result)
	})

	app.Get("/kompas/categories/:category?/:subcategories?", func(c *fiber.Ctx) error {
		if subcategories == "" {
			return c.SendString("You search for " + c.Params("category"))
		}

		return c.SendString("You search for " + c.Params("subcategories") + " in category " + c.Params("category"))
	})

	app.Get("/kompas/news/:url?", func(c *fiber.Ctx) error {
		if c.Params("url") == "" {
			errorText := "Please specify url\n\nExample: /kompas/news/https://otomotif.kompas.com/read/2024/01/06/180829115/pindad-bikin-prototipe-motor-listrik-ev-scooter-daya-jelajah-100-km"


			return c.SendString(errorText)
		}
		return c.SendString("You search for " + c.Params("url"))
	})

	// Start the app on the specified port
	app.Listen(":" + port)
}
