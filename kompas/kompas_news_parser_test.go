package kompas_test

import (
	"testing"

	kompas "github.com/dhupee/Indonesia-News-Aggregator/kompas"
)

func TestKompasGetContent(t *testing.T) {
	url := "https://otomotif.kompas.com/read/2024/01/06/180829115/pindad-bikin-prototipe-motor-listrik-ev-scooter-daya-jelajah-100-km"

	kompasNews, err := kompas.KompasGetData(url, &kompas.KompasNewsStruct{})

	if err != nil {
		t.Error(err)
	}

	if kompasNews.Title == "" {
		t.Error("Error retrieving title")
	}
	if kompasNews.Author == "" {
		t.Error("Error retrieving author")
	}
	if kompasNews.Editor == "" {
		t.Error("Error retrieving editor")
	}
	if kompasNews.Date == "" {
		t.Error("Error retrieving date")
	}
	if kompasNews.Image == "" {
		t.Error("Error retrieving image")
	}
	if len(kompasNews.Tags) == 0 {
		t.Error("Error retrieving tags")
	}
	if len(kompasNews.Content) == 0 {
		t.Error("Error retrieving content")
	}
}
