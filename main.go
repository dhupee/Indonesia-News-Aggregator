package main

import (
	"log"
	"os"

	kompas "github.com/dhupee/Indonesia-News-Aggregator/kompas"
	handlers "github.com/dhupee/Indonesia-News-Aggregator/handlers"

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

	// Define the route handler for the root path and non-source-specific routes
	app.Get("/", handlers.RootHandler)

	// Define the route handlers for Kompas endpoints
	app.Get("/kompas/search/:keyword?", KompasSearchHandler)
	app.Get("/kompas/categories/:category?/:subcategories?", KompasCategoriesHandler)
	app.Get("/kompas/news/", KompasNewsHandler)

	// Start the app on the specified port
	app.Listen(":" + port)
}

func KompasSearchHandler(c *fiber.Ctx) error {
	keyword := c.Params("keyword")
	if keyword == "" {
		return c.SendString("Please specify keyword")
	}

	result := kompas.Search(keyword)
	return c.SendString("You search for " + result)
}

func KompasCategoriesHandler(c *fiber.Ctx) error {
	category := c.Params("category")
	subcategories := c.Params("subcategories")

	if subcategories == "" {
		return c.SendString("You search for " + category)
	}

	return c.SendString("You search for " + subcategories + " in category " + category)
}

func KompasNewsHandler(c *fiber.Ctx) error {
	url := c.Get("Source")

	if url == "" {
		return c.SendFile("./kompas/error_text/kompas_news_handler.txt")
	}

	kompasNews := kompas.KompasGetData(url, &kompas.KompasNewsStruct{})

	return c.JSON(kompasNews)
}
