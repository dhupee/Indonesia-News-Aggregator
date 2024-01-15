package kompas

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func KompasIndexHandler(c *fiber.Ctx) error {
	// Header
	page := c.Get("Page")
	date := c.Get("Date")
	category := c.Get("Category")

	var url string

	if category != "" {
		categoryValid := KompasCategoryCheck(category, KompasCategoryList)
		if !categoryValid {
			return c.SendString("Invalid category, please specify one of the following: \n \n" + strings.Join(KompasCategoryList, ", "))
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
	newsIndex, err := KompasGetNewsIndex(url)
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

	kompasNews, err := KompasGetData(url, &KompasNewsStruct{})
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(kompasNews)
}
