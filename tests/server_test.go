package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/database"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/models"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/server"
)

func TestHandleHome(t *testing.T) {
	r := chi.NewRouter()
	s := server.NewServer(nil)
	s.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	expectedBody := "Welcome to Go Web Scraper"
	if rec.Body.String() != expectedBody {
		t.Errorf("expected '%s' got '%s'", expectedBody, rec.Body.String())
	}
}

func TestHandleData_NoDB(t *testing.T) {
	r := chi.NewRouter()
	s := server.NewServer(nil)
	s.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/data", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	// Without a DB connection, the query should fail
	if rec.Code == http.StatusOK {
		t.Errorf("expected an error status, got 200")
	}
}

// Example of a more integrated test with a real DB connection
func TestHandleData_WithDB(t *testing.T) {
	// Setup a test DB (ensure you have a test environment for this!)
	db, err := database.ConnectPostgres("postgres://username:password@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	defer db.Close()

	// Insert dummy data
	_, err = db.Exec("INSERT INTO scraped_data (url, title) VALUES ($1, $2)", "https://testurl.com", "Test URL Title")
	if err != nil {
		t.Fatalf("Failed to insert dummy data: %v", err)
	}

	r := chi.NewRouter()
	s := server.NewServer(db)
	s.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/data", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", rec.Code)
	}

	var results []models.ScrapedData
	if err := json.Unmarshal(rec.Body.Bytes(), &results); err != nil {
		t.Fatalf("unable to parse JSON: %v", err)
	}

	if len(results) == 0 {
		t.Error("expected at least one row in results")
	}
}
