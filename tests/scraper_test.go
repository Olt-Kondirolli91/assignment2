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

	// We expect up to 3 results (if the site times out or fails, it might skip)
	if len(results) == 0 {
		t.Error("Expected results to have at least 1 entry")
	}
}
