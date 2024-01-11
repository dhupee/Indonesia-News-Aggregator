package utils_test

import (
	"testing"

	"github.com/dhupee/Indonesia-News-Aggregator/utils"
)

func TestGetHtml(t *testing.T) {
	url := "https://www.example.com"
	rawHTML, err := utils.GetHtml(url)

	if err != nil {
		t.Error(err)
	}

	if rawHTML == "" {
		t.Error("Empty HTML")
	}
}
