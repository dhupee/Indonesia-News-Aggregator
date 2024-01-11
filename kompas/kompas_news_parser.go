package kompas

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"

	"github.com/PuerkitoBio/goquery"
)

// declare the struct
type KompasNewsStruct struct {
	Url     string
	Title   string
	Author  string
	Editor  string
	Date    string
	Image   string

	Tags    []string
	Content string
}

// * This functions isnt used yet
func newsContentCleanUp(rawNewsContent string) string {
	cleanedNewsContent := []string{}

	pattern := `Copyright 2008 - *`
	regex := regexp.MustCompile(pattern)

	for _, line := range strings.Split(rawNewsContent, ".") {
		if !regex.MatchString(line) {
			cleanedNewsContent = append(cleanedNewsContent, line)
		}
	}

	return strings.Join(cleanedNewsContent, "\n")
}

// KompasGetNewsContent extracts the content from rawHTML based on the specified div tag.
//
// Parameters:
// - rawHtml: the raw HTML string to extract content from.
// - div: the div tag to search for content within.
//
// Return type:
// - string: the extracted content string.
// - error: any error that occurred during the extraction process.
func KompasGetNewsContent(rawHtml string, div string) (string, error) {
	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Extract text content from the specified div tag
	newsContent := doc.Find(div).Text()

	return newsContent, nil
}


// KompasGetTitle retrieves the title from a raw HTML string.
//
// It takes a rawHTML string as a parameter and returns the title string and an error.
//
// Parameters:
// - rawHTML: the raw HTML string to extract the title from.
//
// Returns:
// - string: the title string.
// - error: an error if the title could not be extracted.
func KompasGetTitle(rawHTML string) (string, error) {
	// Load the raw HTML string into a goquery document
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHTML))
	if err != nil {
		log.Println(err)
		return "", err
	}

	// Find the <title> tag and get its text content
	title := doc.Find("title").Text()

	return title, nil
}


// KompasGetMetadata retrieves the specified metadata from a given code block using a regular expression pattern.
//
// Parameters:
//   rawHtml: The code block to search for metadata.
//   pattern: The regular expression pattern to match the metadata.
//
// Returns:
//   string: The value of the matched metadata.
//   error: An error if no match was found.
func KompasGetMetadata(rawHtml string, pattern string) (string, error) {
    re := regexp.MustCompile(pattern)
    matches := re.FindStringSubmatch(rawHtml)
    if len(matches) > 1 {
        return matches[1], nil
    }
    return "", fmt.Errorf("no match found")
}


// KompasGetMetadata2 retrieves metadata from the given raw HTML using the specified pattern.
//
// Parameters:
// - rawHtml: The raw HTML string to extract metadata from.
// - pattern: The CSS selector pattern to locate the metadata element.
//
// Returns:
// - string: The extracted metadata.
// - error: An error if the HTML parsing or pattern matching fails.
func KompasGetMetadata2(rawHtml string, pattern string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		log.Println(err)
	}

	metadata := doc.Find(pattern).AttrOr("content", "")
	if metadata == "" {
		log.Println("no match found on " + pattern)
	}
	// fmt.Println(author)

	return metadata, nil
}



// KompasGetNewsTags retrieves the news tags from the given raw HTML using the specified pattern.
//
// Parameters:
// - rawHtml: the raw HTML string to extract the tags from.
// - pattern: the CSS selector pattern used to locate the tags in the HTML.
//
// Returns:
// - []string: the list of news tags extracted from the HTML.
// - error: an error if any occurred during the extraction process.
func KompasGetNewsTags(rawHtml string, pattern string) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		return nil, err
	}

	var tag_list []string

	tags := doc.Find(pattern).AttrOr("content", "")
	tag_list = strings.Split(tags, ", ")

	return tag_list, nil
}


// kompasGetImageUrl extracts the URL of an image from the given raw HTML using goquery.
//
// rawHtml: The raw HTML string from which to extract the image URL.
// Returns the extracted image URL as a string.
func KompasGetImageUrl(rawHtml string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		// Handle the error
		log.Println("Error parsing HTML:", err)
		return "", err
	}

	var imageUrl string
	doc.Find("link[rel='preload'][as='image']").Each(func(i int, s *goquery.Selection) {
		imageUrl, _ = s.Attr("href")
	})

	return imageUrl, nil
}



// KompasGetData retrieves data from the given URL and populates the KompasNewsStruct with the extracted information.
//
// Parameters:
// - url: a string representing the URL to retrieve the data from.
// - kompasNews: a pointer to a KompasNewsStruct where the extracted information will be stored.
//
// Return:
// - KompasNewsStruct: the populated KompasNewsStruct containing the extracted data.
func KompasGetData(url string, kompasNews *KompasNewsStruct) KompasNewsStruct {
	// get the raw HTML
	rawHTML, err := utils.GetHtml(url)
	if err != nil {
		log.Println(err)
		return KompasNewsStruct{}, err
	}

	title, err := KompasGetTitle(rawHTML)
	if err != nil {
		log.Println(err)
		title = ""
	}

	// author, err := kompasGetMetadata(rawHTML, `"content_author":\s+"([^"]+)`)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	author, err := KompasGetMetadata2(rawHTML, "meta[name='content_author']")
	if err != nil {
		log.Println(err)
		author = ""
	}

	editor, err := KompasGetMetadata2(rawHTML, "meta[name='content_editor']")
	if err != nil {
		log.Println(err)
		editor = ""
	}

	date, err := KompasGetMetadata2(rawHTML, "meta[name='content_PublishedDate']")
	if err != nil {
		log.Println(err)
		date = ""
	}

	image, err := KompasGetImageUrl(rawHTML)
	if err != nil {
		log.Println(err)
		image = ""
	}

	newsContent, err := KompasGetNewsContent(rawHTML, `read__content`)
	if err != nil {
		log.Println(err)
		newsContent = ""
	}

	newsTags, err := KompasGetNewsTags(rawHTML, "meta[name='content_tags']")
	if err != nil {
		log.Println(err)
		newsTags = []string{}
	}

	// assign values to the struct fields
	kompasNews.Url = url
	kompasNews.Title = title
	kompasNews.Author = author
	kompasNews.Editor = editor
	kompasNews.Date = date
	kompasNews.Image = image

	kompasNews.Content = newsContent
	kompasNews.Tags = newsTags

	return *kompasNews, err
}
