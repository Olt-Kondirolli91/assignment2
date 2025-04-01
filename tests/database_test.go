package tests

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/database"
)

func TestConnectPostgres(t *testing.T) {
	db, err := database.ConnectPostgres("postgres://username:password@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		t.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	// Check if we can run a basic query
	var result int
	if err := db.QueryRow("SELECT 1").Scan(&result); err != nil {
		t.Fatalf("Basic query failed: %v", err)
	}
	if result != 1 {
		t.Errorf("Expected 1, got %d", result)
	}
}
