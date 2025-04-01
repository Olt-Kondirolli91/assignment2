package tests

import (
	"testing"

	"github.com/Olt-Kondirolli91/go-web-scraper/internal/scraper"
)

func TestScrapeSites_Multiple(t *testing.T) {
	urls := []string{
		"https://example.com", 
		"https://example.org", 
		"https://example.net",
	}

	results, err := scraper.ScrapeSites(urls)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(results) != len(urls) {
		t.Errorf("Expected %d results, got %d", len(urls), len(results))
	}
}
