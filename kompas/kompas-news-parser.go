package kompas

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

//TODO: add functions to append it to an object to be cleaned up
func ExtractContent(rawHTML string, div string) string {
	// Create a new tokenizer from the raw HTML
	tokenizer := html.NewTokenizer(strings.NewReader(rawHTML))

	// Variable to track whether the current token is within a <p> tag
	var inPTag bool

	// Iterate through the tokens
	for {
		tokenType := tokenizer.Next()

		switch tokenType {
		case html.ErrorToken:
			// End of the document
			fmt.Println("End of document")
			return ""

		case html.StartTagToken:
			token := tokenizer.Token()

			// Check if the current tag is a <p> tag
			if token.Data == "p" {
				inPTag = true
			}

		case html.EndTagToken:
			token := tokenizer.Token()

			// Check if the current tag is the closing tag of a <p> tag
			if inPTag && token.Data == "p" {
				inPTag = false
			}

		case html.TextToken:
			// If you want to extract text content within <p> tags, add your logic here
			if inPTag {
				text := strings.TrimSpace(tokenizer.Token().Data)
				if text != "" {
					fmt.Println("Text within <p> tag:", text)
				}
			}
		}
	}
}
