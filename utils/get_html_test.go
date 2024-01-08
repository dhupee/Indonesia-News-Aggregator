package utils

import (
	"testing"
)

func TestGetHtml(t *testing.T) {
	url := "https://www.example.com"
	rawHTML := GetHtml(url)
	if rawHTML == "" {
		t.Error("Error")
	}
}
