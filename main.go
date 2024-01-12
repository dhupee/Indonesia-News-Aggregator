package main

import (
	"log"
	"os"
	"regexp"
	"strings"

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
	// Header
	page := c.Get("Page")
	date := c.Get("Date")
	category := c.Get("Category")

	var url string

	if category != "" {
		categoryValid := kompas.KompasCategoryCheck(category, kompas.KompasCategoryList)
		if !categoryValid {
			return c.SendString("Invalid category, please specify one of the following: \n \n" + strings.Join(kompas.KompasCategoryList, ", "))
		}
	}

	//make sure the date is YYYY-MM-DD
	if date != "" {
		if !regexp.MustCompile(`^([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))$`).MatchString(date) {
			return c.SendString("Invalid date, please use YYYY-MM-DD format in your 'Date' header")
		}
	}

	switch {
	case category != "" && page != "" && date != "":
		url = "https://indeks.kompas.com/?site=" + category + "&page=" + page + "&date=" + date
	case page != "" && date != "":
		url = "https://indeks.kompas.com/?page=" + page + "&date=" + date
	case category != "" && date != "":
		url = "https://indeks.kompas.com/?site=" + category + "&date=" + date
	case category != "" && page != "":
		url = "https://indeks.kompas.com/?site=" + category + "&page=" + page
	case date != "":
		url = "https://indeks.kompas.com/?date=" + date
	case page != "":
		url = "https://indeks.kompas.com/?page=" + page
	case category != "":
		url = "https://indeks.kompas.com/?site=" + category
	default:
		url = "https://indeks.kompas.com/"
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
