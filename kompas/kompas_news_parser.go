package kompas

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"

	"golang.org/x/net/html"
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
	Content []string
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

func KompasExtractContentFromDiv(rawHTML string, div string) []string {
	tokenizer := html.NewTokenizer(strings.NewReader(rawHTML))

	newsContent := []string{}

	var inPTag bool

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			// Reached the end of the document, return the extracted content
			return newsContent

		case html.StartTagToken:
			token := tokenizer.Token()

			if token.Data == "p" {
				inPTag = true
			}

		case html.EndTagToken:
			token := tokenizer.Token()

			if inPTag && token.Data == "p" {
				inPTag = false
			}

		case html.TextToken:
			if inPTag {
				text := strings.TrimSpace(tokenizer.Token().Data)
				if text != "" {
					log.Println("Extracting text:", text)
					newsContent = append(newsContent, text)
				}
			}
		}
	}
}

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

// kompasExtractContentFromScriptTag extracts the content from a script tag in a code block using a regular expression pattern.
//
// It takes two parameters:
// - codeBlock: a string representing the code block containing the script tag.
// - pattern: a string representing the regular expression pattern to match the script tag content.
//
// It returns a string representing the content of the script tag if a match is found, and an error otherwise.
func kompasExtractContentFromScriptTag(codeBlock string, pattern string) (string, error) {
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

// kompasExtractContentTags extracts content tags from a given code block using a regular expression pattern.
//
// Parameters:
//   - codeBlock: The code block from which to extract the tags.
//   - pattern: The regular expression pattern used to match the tags.
//
// Returns:
//   - []string: A slice of strings containing the extracted tags.
//   - error: An error if any occurred during the extraction process.
func kompasExtractContentTags(codeBlock string, pattern string) ([]string, error) {
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

// kompasExtractImageUrl extracts the URL of an image from raw HTML.
//
// rawHtml: the raw HTML string to search for the image URL.
// returns: the URL of the image, or an empty string if no URL is found.
func kompasExtractImageUrl(rawHtml string) string {
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


// KompasGetContent retrieves content from a given URL and populates a KompasNewsStruct.
//
// Parameters:
// - url: a string representing the URL to fetch the content from.
// - kompasNews: a pointer to a KompasNewsStruct object that will be populated with the retrieved content.
//
// Returns:
// - KompasNewsStruct: the populated KompasNewsStruct object.
func KompasGetContent(url string, kompasNews *KompasNewsStruct) KompasNewsStruct {
	// get the raw HTML
	rawHTML := utils.GetHtml(url)

	title := KompasGetTitle(rawHTML)

	author, err := kompasExtractContentFromScriptTag(rawHTML, `"content_author":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
	}

	editor, err := kompasExtractContentFromScriptTag(rawHTML, `"content_editor":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
	}

	date, err := kompasExtractContentFromScriptTag(rawHTML, `"content_PublishedDate":\s+"([^"]+)`)
	if err != nil {
		log.Fatal(err)
	}

	image := kompasExtractImageUrl(rawHTML)

	newsContent := KompasExtractContentFromDiv(rawHTML, "read__content")
	newsTags, err := kompasExtractContentTags(rawHTML, `"content_tags":\s+"([^"]+)`)
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
