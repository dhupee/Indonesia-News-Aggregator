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
type kompasNewsStruct struct {
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

func KompasExtractContentFromDiv(rawHTML string, div string) [] string {
	tokenizer := html.NewTokenizer(strings.NewReader(rawHTML))

	newsContent := []string{}

	var inPTag bool

	// stopPattern := regexp.MustCompile("Copyright 2008 - 2023 PT. Kompas Cyber Media (Kompas Gramedia Digital Group).")

	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			log.Println("End of document")
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

func kompasExtractContentFromScriptTag(codeBlock string, pattern string) (string, error) {
	// Compile the regular expression
	re := regexp.MustCompile(pattern)

	// Find the matches
	matches := re.FindStringSubmatch(codeBlock)

	// Check if a match was found
	if len(matches) > 1 {
		contentTitle := matches[1]
		return contentTitle, nil
	}

	return "", fmt.Errorf("no match found")
}

func KompasGetContent(url string, kompasNews *kompasNewsStruct) *kompasNewsStruct {
	rawHTML := utils.GetHtml(url)

	title, err := kompasExtractContentFromScriptTag(rawHTML, `content_title":\s*"([^"]+)"`)
	if err != nil {
		log.Fatal(err)
	}
	newsContent := KompasExtractContentFromDiv(rawHTML, "read__content")

	// assign values to the struct fields
	kompasNews.Url = url
	kompasNews.Title = title
	kompasNews.Content = newsContent

	return kompasNews
}
