package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/models"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/scraper"
)

// Server holds references to shared resources, such as the database.
type Server struct {
	DB *sql.DB
}

// NewServer creates a new Server instance with the provided db connection. 
func NewServer(db *sql.DB) *Server {
	return &Server{DB: db}
}

// SetupRoutes registers all HTTP endpoints.
func (s *Server) SetupRoutes(r *chi.Mux) {
	r.Get("/", s.handleHome)
	r.Get("/scrape", s.handleScrape)
	r.Get("/data", s.handleGetData)
}

func (s *Server) handleHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to Go Web Scraper"))
}

func (s *Server) handleScrape(w http.ResponseWriter, r *http.Request) {
	urls := []string{
		"https://example.org",
		"https://example.com",
	}

	results, err := scraper.ScrapeSites(urls)
	if err != nil {
		http.Error(w, "Scraping failed", http.StatusInternalServerError)
		return
	}

	for _, res := range results {
		_, err := s.DB.Exec("INSERT INTO scraped_data (url, title) VALUES ($1, $2)", res.URL, res.Title)
		if err != nil {
			continue
		}
	}

	w.Write([]byte("Scraping done and data saved!"))
}

func (s *Server) handleGetData(w http.ResponseWriter, r *http.Request) {
	rows, err := s.DB.Query("SELECT id, url, title FROM scraped_data")
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []models.ScrapedData
	for rows.Next() {
		var d models.ScrapedData
		rows.Scan(&d.ID, &d.URL, &d.Title)
		results = append(results, d)
	}

	json.NewEncoder(w).Encode(results)
}
