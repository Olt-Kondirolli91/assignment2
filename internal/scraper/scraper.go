package scraper

import (
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"

	"github.com/Olt-Kondirolli91/go-web-scraper/internal/models"
)

// ScrapeSites concurrently fetches each URL in urls, 
// extracts the <title> element, and returns a slice of ScrapedData.
func ScrapeSites(urls []string) ([]models.ScrapedData, error) {
	var wg sync.WaitGroup
	resultsChan := make(chan models.ScrapedData, len(urls))

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			title, err := scrapeTitle(url)
			if err == nil {
				resultsChan <- models.ScrapedData{URL: url, Title: title}
			}
		}(u)
	}

	wg.Wait()
	close(resultsChan)

	var results []models.ScrapedData
	for r := range resultsChan {
		results = append(results, r)
	}
	return results, nil
}

// scrapeTitle fetches the URL, parses the HTML,
// and returns the <title> element text.
func scrapeTitle(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return "", err
	}

	var title string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			title = strings.TrimSpace(n.FirstChild.Data)
			return
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return title, nil
}
