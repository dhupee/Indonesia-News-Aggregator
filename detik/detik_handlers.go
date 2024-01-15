package detik

import (
	"regexp"

	"github.com/gofiber/fiber/v2"
)

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
