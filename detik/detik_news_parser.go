package detik

import (
	"fmt"
	"log"
	"strings"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"

	"github.com/PuerkitoBio/goquery"

)

type DetikNewsStruct struct {
	Url    string
	Title  string
	Author string
	Editor string
	Date   string
	Image  string

	Tags    []string
	Content string
}


func DetikGetTitle(content string, pattern string) (string, error) {
    doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
    if err != nil {
        return "", fmt.Errorf("could not create document: %v", err)
    }
    newsTitle := doc.Find(pattern).Text()

    return newsTitle, nil
}

func DetikGetMetadata(content string, pattern string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return "", fmt.Errorf("could not create document: %v", err)
	}
	newsMetadata := doc.Find(pattern).AttrOr("content", "")

	return newsMetadata, nil
}

func DetikGetTags(content string, pattern string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		return []string{}, fmt.Errorf("could not create document: %v", err)
	}
	tags := doc.Find(pattern).AttrOr("content", "")
	tag_list := strings.Split(tags, ",")

	return tag_list, nil
}

func DetikGetContent(rawHtml string, pattern string) (string, error) {
	var newsContent string

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		log.Println("could not create document: %v", err)
	}

	doc.Find(pattern).Each(func(i int, s *goquery.Selection) {
		newsContent = newsContent + s.Text()
	})

	return newsContent, err
}

func DetikGetData(url string, detikNews *DetikNewsStruct) (DetikNewsStruct, error) {

	// get the raw HTML
	rawHtml, err := utils.GetHtml(url)
	if err != nil {
		log.Println(err)
		return DetikNewsStruct{}, err
	}

	// get the title
	title, err := DetikGetTitle(rawHtml,"title")
	if err != nil {
		log.Println(err)
		title = ""
	}

	// get the author
	author, err := DetikGetMetadata(rawHtml, `meta[name="author"]`)
	if err != nil {
		log.Println(err)
		author = ""
	}

	// get the editor
	// use empty strings since no info
	editor := ""

	// get the date
	date, err := DetikGetMetadata(rawHtml, `meta[name="publishdate"]`)
	if err != nil {
		log.Println(err)
		date = ""
	}

	// get the image url
	imageUrl, err := DetikGetMetadata(rawHtml, `meta[name="thumbnailUrl"]`)
	if err != nil {
		log.Println(err)
		imageUrl = ""
	}

	// get the tags
	tags, err := DetikGetTags(rawHtml, `meta[name="keywords"]`)
	if err != nil {
		log.Println(err)
		tags = []string{}
	}

	// get the content
	content, err := DetikGetContent(rawHtml, ".detail__body-text p")
	if err != nil {
		log.Println(err)
		content = ""
	}

	// assign the values to the DetikNewsStruct
	detikNews.Url = url
	detikNews.Title = title
	detikNews.Author = author
	detikNews.Editor = editor
	detikNews.Date = date
	detikNews.Image = imageUrl
	detikNews.Tags = tags
	detikNews.Content = content

	return *detikNews, nil
}

