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
		log.Println("Failed to load .env")
		log.Println("Using default environment variables")
	}

	// Get the port from the environment variables
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	// Create a new Fiber app
	app := fiber.New(fiber.Config{
		ServerHeader: "Indonesia-News-Aggregator",
	})

	v1 := app.Group("/v1") // add v1 grouping to manage if needed

	// Define the route handler for the root path and non-source-specific routes
	v1.Get("/", RootHandler)
	// v1.Get("/search", SearchHandler)

	// Define the route handlers for Kompas endpoints
	v1.Get("/kompas/index", KompasIndexHandler)
	v1.Get("/kompas/news", KompasNewsHandler)

	// Start the app on the specified port
	log.Fatal(app.Listen(":"+port))
}

func RootHandler(c *fiber.Ctx) error {
	return c.SendFile("./assets/welcome.txt")
}

// TODO: this one next, maybe just use universal solution
// TODO: W.I.P
// func SearchHandler(c *fiber.Ctx) error {
// 	keyword := c.Get("keyword")
// 	if keyword == "" {
// 		return c.SendString("Please specify keyword")
// 	}

// 	result := kompas.Search(keyword)
// 	return c.SendString("You search for " + result)
// }

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
			return c.SendString("Invalid date\n\nPlease use YYYY-MM-DD format in your 'Date' header")
		}
	}

	// make sure the page is a number
	if page != "" {
		if !regexp.MustCompile(`^[0-9]+$`).MatchString(page) {
			return c.SendString("Invalid page\n\nPlease use a number in your 'Page' header")
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
			return c.SendFile("./assets/kompas/kompas_news_handler_error.txt")
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
