package detik

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func DetikIndexHandler(c *fiber.Ctx) error {
	// Header
	page := c.Get("Page")
	date := c.Get("Date")
	category := c.Get("Category")

	var url string

	if category != "" {
		categoryValid := DetikCategoryCheck(category, DetikCategoryList)
		if !categoryValid {
			return c.SendString("Invalid category, please specify one of the following: \n \n" + strings.Join(DetikCategoryList, ", "))
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
	case category == "news":
		url = "https://news.detik.com/indeks/"
	case category == "edu":
		url = "https://www.detik.com/edu/indeks/"
	case category == "finance":
		url = "https://finance.detik.com/indeks/"
	case category == "hot":
		url = "https://hot.detik.com/indeks/"
	case category == "inet":
		url = "https://inet.detik.com/indeks/"
	case category == "oto":
		url = "https://oto.detik.com/indeks/"
	case category == "travel":
		url = "https://travel.detik.com/indeks/"
	case category == "sepakbola":
		url = "https://sport.detik.com/sepakbola/indeks"
	case category == "food":
		url = "https://food.detik.com/indeks/"
	case category == "health":
		url = "https://health.detik.com/indeks/"
	case category == "wolipop":
		url = "https://wolipop.detik.com/indeks/"
	case category == "jatim":
		url = "https://www.detik.com/jatim/indeks/"
	case category == "jateng":
		url = "https://www.detik.com/jateng/indeks/"
	case category == "jabar":
		url = "https://www.detik.com/jabar/indeks/"
	case category == "sulsel":
		url = "https://www.detik.com/sulsel/indeks/"
	case category == "sumut":
		url = "https://www.detik.com/sumut/indeks/"
	case category == "bali":
		url = "https://www.detik.com/bali/indeks/"
	case category == "hikmah":
		url = "https://www.detik.com/hikmah/indeks/"
	case category == "sumbagsel":
		url = "https://www.detik.com/sumbagsel/indeks/"
	case category == "properti":
		url = "https://www.detik.com/properti/indeks/"
	case category == "jogja":
		url = "https://www.detik.com/jogja/indeks/"
	default:
		url = "https://news.detik.com/indeks/"
	}

	if page != "" && date != "" {
		// extract the DD, MM, and YYYY from the date
		dateParts := strings.Split(date, "-") 
		day := dateParts[2]
		month := dateParts[1]
		year := dateParts[0]

		url = url + page + "?date=" + month + "/" + day + "/" + year
	} else if page != "" {
		url = url + page
	} else if date != "" {
		// extract the DD, MM, and YYYY from the date
		dateParts := strings.Split(date, "-")
		day := dateParts[2]
		month := dateParts[1]
		year := dateParts[0]

		url = url + "?date=" + month + "/" + day + "/" + year
	}

	// get news index
	newsIndex, err := DetikGetNewsIndex(url)
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(newsIndex)
}

func DetikNewsHandler(c *fiber.Ctx) error {
	url := c.Get("Source")

	subDomainRegex := regexp.MustCompile(`^https?://(.+\.)*detik\.com`)
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

	detikNews, err := DetikGetData(url, &DetikNewsStruct{})
	if err != nil {
		return c.SendString(err.Error())
	}
	return c.JSON(detikNews)
}
