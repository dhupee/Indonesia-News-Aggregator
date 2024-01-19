package detik

import (
	"log"
	"strings"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"

	"github.com/PuerkitoBio/goquery"
)

type Article struct{
	Title string
	Url string
	Date string
	Image string
}

func DetikGetNewsIndex(url string) ([]Article, error) {
	// get the HTML content
	rawHtml, err := utils.GetHtml(url)
	if err != nil {
		log.Println(err)
	}

	DetikNewsIndex := []Article{}

	// get the news index
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		log.Println("could not parse the content: %v", err)
	}
	articles := doc.Find(".grid-row.list-content article.list-content__item")

	// Loop through each article and extract the desired information
	articles.Each(func(i int, article *goquery.Selection) {
		// Extract the article title
		title := article.Find(".media__title a").Text()

		// Extract the article URL
		url, _ := article.Find(".media__title a").Attr("href")

		// Extract the article date
		date := article.Find(".media__date span").AttrOr("title", "")

		// Extract the image URL
		imageURL, _ := article.Find(".media__image img").Attr("src")

		// add values to the index
		DetikNewsIndex = append(DetikNewsIndex, Article{
			Title: title,
			Url:   url,
			Date:  date,
			Image: imageURL,
		})
	})

	return DetikNewsIndex, nil
}

func DetikCategoryCheck(category string, categoryList []string) bool {
	var results bool
	if utils.IsInSlice(category, categoryList){
		results = true
	} else{
		results = false
	}
	return results
}
