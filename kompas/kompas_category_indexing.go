package kompas

import (
	// "fmt"
	// "time"
	"log"
	"strings"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"

	"github.com/PuerkitoBio/goquery"
)

// declare the struct
type Article struct{
	Title string
	Url string
	Date string
	Image string
}

var KompasCategoryList = []string{
	"all",
	"nasional",
	"regional",
	"megapolitan",
	"global",
	"tren",
	"health",
	"food",
	"edukasi",
	"money",
	"properti",
	"bola",
	"travel",
	"otomotif",
	"sains",
	"hype",
	"jeo",
	"skola",
	"stori",
	"konsultasihukum",
	"wiken",
	"headline",
	"terpopuler",
	"sorotan",
	"topik",
	"advertorial",
}

func KompasGetNewsIndex(url string) ([]Article, error) {
	rawHtml, err := utils.GetHtml(url)
	if err != nil {
		log.Println(err)
	}

	newsIndex := []Article{}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		log.Println(err)
		return newsIndex, err
	}

	// Extract information from the HTML
	doc.Find(".article__list").Each(func(i int, s *goquery.Selection) {
		// Get article title
		title := s.Find(".article__title").Text()

		// Get article URL
		url, _ := s.Find(".article__link").Attr("href")

		// Get article date
		date := s.Find(".article__date").Text()

		image := s.Find(".article__asset").Find("img").AttrOr("src", "")

		newsIndex = append(newsIndex, Article{title, url, date, image})
	})

	return newsIndex, nil
}

func KompasCategoryCheck(category string, categoryList []string) bool {
	var results bool
	if utils.IsInSlice(category, categoryList){
		results = true
	} else{
		results = false
	}
	return results
}
