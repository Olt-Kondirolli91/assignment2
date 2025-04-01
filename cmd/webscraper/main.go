package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/database"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/server"
)

func main() {
	// Adjust the DSN to match your PostgreSQL credentials.
	dsn := "postgres://postgres:test@localhost:5432/golang?sslmode=disable"

	db, err := database.ConnectPostgres(dsn)
	if err != nil {
		log.Println("Could not connect to DB:", err)
		// You can continue without a DB or exit. Here, let's exit.
		return
	}
	defer db.Close()

	r := chi.NewRouter()

	s := server.NewServer(db)
	s.SetupRoutes(r)

	log.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Server failed:", err)
	}
}
