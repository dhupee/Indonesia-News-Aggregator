package kompas

import (
	"io"
	"log"
	"net/http"
)

func Search(keyword string) string {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://search.kompas.com/search/?q="+keyword+"", nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	
	result := string(bodyText)
	return result
}
