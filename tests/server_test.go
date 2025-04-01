package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/database"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/models"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/server"
)

// TestHandleHome ensures the root route returns the expected welcome message.
func TestHandleHome(t *testing.T) {
	r := chi.NewRouter()
	srv := server.NewServer(nil) // nil DB is okay here because / won't query DB
	srv.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rec.Code)
	}

	expected := "Welcome to Go Web Scraper"
	if rec.Body.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, rec.Body.String())
	}
}

// TestHandleData_NoDB checks that the endpoint returns an error if DB is nil.
func TestHandleData_NoDB(t *testing.T) {
	r := chi.NewRouter()
	srv := server.NewServer(nil) // We pass nil to simulate no DB
	srv.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/data", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	// We expect 500 because the DB is nil
	if rec.Code != http.StatusInternalServerError {
		t.Fatalf("Expected status 500, got %d", rec.Code)
	}
}

// TestHandleData_WithDB is an integration test requiring a real test DB.
func TestHandleData_WithDB(t *testing.T) {
	// Change this to point to your test DB
	testDSN := "postgres://postgres:test@localhost:5432/golang?sslmode=disable"

	db, err := database.ConnectPostgres(testDSN)
	if err != nil {
		t.Fatalf("Failed to connect to test DB: %v", err)
	}
	defer db.Close()

	// Insert dummy data
	_, err = db.Exec("INSERT INTO scraped_data (url, title) VALUES ($1, $2)", "https://test.com", "Test Title")
	if err != nil {
		t.Fatalf("Failed to insert dummy data: %v", err)
	}

	r := chi.NewRouter()
	srv := server.NewServer(db)
	srv.SetupRoutes(r)

	req, _ := http.NewRequest("GET", "/data", nil)
	rec := httptest.NewRecorder()
	r.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("Expected status 200, got %d", rec.Code)
	}

	var results []models.ScrapedData
	if err := json.Unmarshal(rec.Body.Bytes(), &results); err != nil {
		t.Fatalf("Failed to parse JSON: %v", err)
	}

	if len(results) == 0 {
		t.Error("Expected at least one row in results")
	}
}
