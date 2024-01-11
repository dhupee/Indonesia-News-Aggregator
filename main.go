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
	app.Get("/kompas/index/", KompasIndexHandler)
	app.Get("/kompas/news/", KompasNewsHandler)

	// Start the app on the specified port
	app.Listen(":" + port)
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendFile("./assets/welcome.txt")
}

// TODO: this one next, maybe just use universal solution
func KompasSearchHandler(c *fiber.Ctx) error {
	keyword := c.Get("keyword")
	if keyword == "" {
		return c.SendString("Please specify keyword")
	}

	result := kompas.Search(keyword)
	return c.SendString("You search for " + result)
}

func KompasIndexHandler(c *fiber.Ctx) error {
	category := c.Get("category")
	page := c.Get("page")
	date := c.Get("date")

	var url string

	categoryValid := kompas.KompasCategoryCheck(category, kompas.KompasCategoryList)
	if !categoryValid {
		return c.SendString("Invalid category")
	}

	// // wait 3 seconds
	// time.Sleep(3 * time.Second)

	switch {
	case category != "":
		url = "https://indeks.kompas.com/?site=" + category
	case page != "":
		url = "https://indeks.kompas.com/?page=" + page
	case date != "":
		url = "https://indeks.kompas.com/?date=" + date
	case category != "" && page != "":
		url = "https://indeks.kompas.com/?site=" + category + "&page=" + page
	case category != "" && date != "":
		url = "https://indeks.kompas.com/?site=" + category + "&date=" + date
	case page != "" && date != "":
		url = "https://indeks.kompas.com/?page=" + page + "&date=" + date
	case category != "" && page != "" && date != "":
		url = "https://indeks.kompas.com/?site=" + category + "&page=" + page + "&date=" + date
	default:
		url = "https://indeks.kompas.com/?site=tekno"
	}

	// get news index
	newsIndex, err := kompas.KompasGetNewsIndex(url)
	if err != nil {
		return c.SendString(err.Error())
	}

	return c.JSON(newsIndex)
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

	kompasNews, err := kompas.KompasGetData(url, &kompas.KompasNewsStruct{})
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(kompasNews)
}
