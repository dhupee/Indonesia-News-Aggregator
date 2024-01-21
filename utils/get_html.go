package utils

import (
	"io"
	"log"
	"net/http"

	"github.com/playwright-community/playwright-go"
)

// GetHtmlSimple retrieves the HTML content from the specified URL using simple HTTP requests.
//
// Parameters:
// - url: a string representing the URL to be fetched.
//
// Returns:
// - string: the HTML content of the URL.
// - error: an error if any occurred during the request.
func GetHtmlSimple(url string) (string, error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return "", err
	}

	rawHTML := string(bodyText)

	return rawHTML, nil
}

// GetHtml retrieves the HTML content of a web page given its URL using Playwright to scrape the web page.
//
// url: The URL of the web page.
// string: The HTML content of the web page.
// error: An error if something goes wrong during the process.
func GetHtml(url string) (string, error) {
	// start the browser, if not installed then install
	pw, err := playwright.Run()
	if err != nil {
		log.Println("could not start playwright, installing...")
		if err = playwright.Install(); err != nil {
			log.Fatal("could not install playwright: %v", err)
		}
	}

	// launch the browser
	browser, err := pw.Chromium.Launch()
	if err != nil {
		log.Println("could not launch browser: %v", err)
	}

	// create a new page
	page, err := browser.NewPage()
	if err != nil {
		log.Println("could not create page: %v", err)
	}

	// goto the url
	if _, err = page.Goto(url); err != nil {
		log.Println("could not goto url: %v", err)
	}

	// // scroll all the way to the bottom
	// for i := 0; i < 10; i++ {
	// 	_, err = page.Evaluate(`window.scrollTo(0, document.body.scrollHeight)`)
	// 	if err != nil {
	// 		log.Println("could not scroll: %v", err)
	// 	}
	// }

	// get the content
	rawHtml, err := page.Content()
	if err != nil {
		log.Println("could not get the content: %v", err)
	}

	// close the browser
	err = browser.Close()
	if err != nil {
		log.Println("could not close the browser: %v", err)
	}

	return rawHtml, nil
}
