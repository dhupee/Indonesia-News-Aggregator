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
// - []string: an array of extracted content strings.
func KompasGetNewsContent(rawHtml string, div string) string {
	// Parse the HTML
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(rawHtml))
	if err != nil {
		log.Fatal(err)
	}

	// Extract text content from the "read__content" div
	newsContent := doc.Find(".read__content").Text()

	return newsContent
}

// KompasGetTitle captures the string inside the <title> tag in the given raw HTML.
//
// Parameters:
// - rawHTML: The raw HTML string.
//
// Returns:
// - string: The captured string inside the <title> tag, or an empty string if there is no match.
func KompasGetTitle(rawHTML string) string {
	// capture string inside <title>
	pattern := `<title>(.*?)</title>`

	// compile the regular expression
	re := regexp.MustCompile(pattern)

	// find the first match of the pattern in the rawHTML
	match := re.FindStringSubmatch(rawHTML)

	// if there is a match, return the captured string
	if len(match) > 1 {
		return match[1]
	}

	// if there is no match, return an empty string
	return ""
}


// kompasGetMetadata retrieves the specified metadata from a given code block using a regular expression pattern.
//
// Parameters:
// - codeBlock: The code block to search for metadata.
// - pattern: The regular expression pattern to match the metadata.
//
// Returns:
// - string: The value of the matched metadata.
// - error: An error if no match was found.
func kompasGetMetadata(codeBlock string, pattern string) (string, error) {
	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the matches
	matches := re.FindStringSubmatch(codeBlock)

	// Check if a match was found
	if len(matches) > 1 {
		contentValue := matches[1]
		return contentValue, nil
	}

	return "", fmt.Errorf("no match found")
}


// kompasGetNewsTags returns the tags extracted from a given code block using a regular expression pattern.
//
// Parameters:
// - codeBlock: the code block to search for matches.
// - pattern: the regular expression pattern to match against the code block.
//
// Returns:
// - tags: a slice of strings representing the matched tags.
// - error: an error if any occurred during the matching process.
func kompasGetNewsTags(codeBlock string, pattern string) ([]string, error) {
	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	tags := []string{}

	// Find the matches
	matches := re.FindStringSubmatch(codeBlock)

	// convert the matches to a slice of strings
	// "motor listrik, Pindad, motor, motor listrik pindad, Pindad EV-Scooter"
	if len(matches) > 1 {
		tags = strings.Split(matches[1], ", ")
	}

	return tags, nil
}


// kompasGetImageUrl extracts the URL of an image from the given raw HTML using a regular expression pattern.
//
// rawHtml: The raw HTML string from which to extract the image URL.
// Returns the extracted image URL as a string.
func kompasGetImageUrl(rawHtml string) string {
	// Define the regular expression pattern
	pattern := `<link rel="preload" as="image" href="([^"]+)"`

	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the match in the HTML
	match := re.FindStringSubmatch(rawHtml)

	// Extract the URL from the match
	if len(match) > 1 {
		return match[1]
	}

	return ""
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
	rawHTML := utils.GetHtml(url)

	title := KompasGetTitle(rawHTML)

	author, err := kompasGetMetadata(rawHTML, `"content_author":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
	}

	editor, err := kompasGetMetadata(rawHTML, `"content_editor":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
	}

	date, err := kompasGetMetadata(rawHTML, `"content_PublishedDate":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
	}

	image := kompasGetImageUrl(rawHTML)

	newsContent := KompasGetNewsContent(rawHTML, `read__content`)
	newsTags, err := kompasGetNewsTags(rawHTML, `"content_tags":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
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

	return *kompasNews
}
