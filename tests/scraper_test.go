package tests

import (
	"testing"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/scraper"
)

func TestScrapeSites(t *testing.T) {
	urls := []string{"https://example.org"}
	results, err := scraper.ScrapeSites(urls)
	if err != nil {
		t.Fatal("Expected no error, got:", err)
	}
	if len(results) == 0 {
		t.Fatal("Expected at least one result")
	}
}
