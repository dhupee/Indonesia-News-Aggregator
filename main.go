package main

import (
	"log"
	"os"
	"regexp"

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

	// Define the route handler for the root path and non-source-specific routes
	app.Get("/", RootHandler)

	// Define the route handlers for Kompas endpoints
	app.Get("/kompas/search/", KompasSearchHandler)
	app.Get("/kompas/categories/", KompasCategoriesHandler)
	app.Get("/kompas/news/", KompasNewsHandler)

	// Start the app on the specified port
	app.Listen(":" + port)
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendFile("./assets/welcome.txt")
}

func KompasSearchHandler(c *fiber.Ctx) error {
	keyword := c.Get("keyword")
	if keyword == "" {
		return c.SendString("Please specify keyword")
	}

	result := kompas.Search(keyword)
	return c.SendString("You search for " + result)
}

// ! FOCUS ON THE MAIN CATEGORY INSTEAD
func KompasCategoriesHandler(c *fiber.Ctx) error {
	category := c.Get("category")
	subcategories := c.Get("subcategories")

	if subcategories == "" {
		return c.SendString("You search for " + category)
	} else {
		return c.SendString("You search for " + subcategories + " in category " + category)
	}

	return c.SendString("You search for " + subcategories + " in category " + category)
}

func KompasNewsHandler(c *fiber.Ctx) error {
	url := c.Get("Source")

	subDomainRegex := regexp.MustCompile(`^https?://(.+\.)*kompas\.com`)
	if !subDomainRegex.MatchString(url) {
		if url == "" {
			return c.SendFile("./kompas/error_text/kompas_news_handler.txt")
		}
		// Reject the URL
		domainRegex := regexp.MustCompile(`^https?://([^/]+)`)
		matches := domainRegex.FindStringSubmatch(url)
		if len(matches) > 1 {
			return c.SendFile("rejected " + matches[1])
		}
	}

	kompasNews := kompas.KompasGetData(url, &kompas.KompasNewsStruct{})
	return c.JSON(kompasNews)
}
