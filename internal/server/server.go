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
	if s.DB == nil {
		http.Error(w, "No DB connection available", http.StatusInternalServerError)
		return
	}

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
			// Log or handle error; skipping for brevity
		}
	}

	w.Write([]byte("Scraping completed and data stored in DB!"))
}

func (s *Server) handleGetData(w http.ResponseWriter, r *http.Request) {
	if s.DB == nil {
		http.Error(w, "No DB connection available", http.StatusInternalServerError)
		return
	}

	rows, err := s.DB.Query("SELECT id, url, title FROM scraped_data")
	if err != nil {
		http.Error(w, "Failed to fetch data", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var allData []models.ScrapedData
	for rows.Next() {
		var data models.ScrapedData
		if scanErr := rows.Scan(&data.ID, &data.URL, &data.Title); scanErr != nil {
			continue
		}
		allData = append(allData, data)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(allData)
}
