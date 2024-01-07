package utils

import (
	"testing"
)

func TestGetHTML(t *testing.T) {
	url := "https://www.example.com"
	rawHTML := GetHTML(url)
	if rawHTML == "" {
		t.Error("Error")
	}
}
