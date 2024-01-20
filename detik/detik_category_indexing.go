package detik

import (
	"log"
	"strings"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"

	"github.com/PuerkitoBio/goquery"
)

// Detik Categories
var DetikCategoryList = []string{
	"news",
	"edu",
	"finance",
	"hot", // this is showbiz
	"inet", // related to internet, gadget and stuff
	"sport",
	"oto", // related to motorcycle and car, all automotives
	"travel",
	"sepakbola",
	"food",
	"health",
	"wolipop", // all things women
	"jatim",
	"jateng",
	"jabar",
	"sulsel",
	"sumut",
	"bali",
	"hikmah", // related to religious, mainly islam
	"sumbagsel", // sumsel, lampung
	"properti",
	"jogja",
}

// Detik News Sub categories
var DetikNewsSubCategoryList = []string{
	"semua-berita", // add special condition for this
	"berita",
	"daerah",
	"internasional",
	"kolom",
	"pro-kontra",
	"foto-news", // TODO: check this later
	"detiktv", // TODO: and this
	"bbc",
	"australiaplus",
	"jawabarat",
	"jawabtengah",
	"jawatimur",
	"suara-pembaca",
	"infografis",
	"investigasi",
	"intermeso",
	"crimestory",
	"pemilu",
	"jabodetabek",
	"hukum",
}

var EduDetikNewsSubCategoryList = []string{

}

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
