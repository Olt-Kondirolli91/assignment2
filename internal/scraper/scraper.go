package scraper

import (
	"net/http"
	"strings"
	"sync"

	"golang.org/x/net/html"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/models"
)

func ScrapeSites(urls []string) ([]models.ScrapedData, error) {
	var wg sync.WaitGroup
	resultChan := make(chan models.ScrapedData, len(urls))

	for _, u := range urls {
		wg.Add(1)
		go func(url string) {
			defer wg.Done()
			title, err := scrapeTitle(url)
			if err == nil {
				resultChan <- models.ScrapedData{URL: url, Title: title}
			}
		}(u)
	}

	wg.Wait()
	close(resultChan)

	var results []models.ScrapedData
	for d := range resultChan {
		results = append(results, d)
	}
	return results, nil
}

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
