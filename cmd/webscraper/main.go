package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/database"
	"github.com/Olt-Kondirolli91/go-web-scraper/internal/server"
)

func main() { // Replace username, password, and dbname with your actual PostgreSQL credentials
	db, err := database.ConnectPostgres("postgres://postgres:test@localhost:5432/golang?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := chi.NewRouter()
	s := server.NewServer(db)
	s.SetupRoutes(r)

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
