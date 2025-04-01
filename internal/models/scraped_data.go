package models

type ScrapedData struct {
	ID    int    `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
}
