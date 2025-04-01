package tests

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/database"
)

// A basic test for your database connection logic.
func TestConnectPostgres(t *testing.T) {
	dsn := "postgres://postgres:test@localhost:5432/golang?sslmode=disable" 
	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		t.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	var result int
	if err := db.QueryRow("SELECT 1").Scan(&result); err != nil {
		t.Fatalf("Failed to run basic query: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
