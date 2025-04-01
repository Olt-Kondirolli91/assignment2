# Go Web Scraper

A simple project demonstrating:
- Go web server with Chi (https://github.com/go-chi/chi)
- Database connectivity with PostgreSQL
- Concurrency using goroutines and channels
- Web scraping using net/http and HTML parsing

-------------------------------------------------------------------

Getting Started

1. Clone the repository
   git clone https://github.com/your-username/go-web-scraper.git
   cd go-web-scraper

2. Install Dependencies
   go mod tidy

3. Update Database Credentials
   - In cmd/webscraper/main.go, replace the DSN in ConnectPostgres with your PostgreSQL credentials.

4. Run the Server
   go run ./cmd/webscraper
   - The server runs at http://localhost:8080

-------------------------------------------------------------------

Endpoints

- GET / — Returns a welcome message.
- GET /scrape — Scrapes predefined URLs, inserts their titles into the scraped_data table.
- GET /data — Returns the scraped data in JSON format.

-------------------------------------------------------------------

Testing

Several tests cover different components:

1. Scraper Tests (tests/scraper_test.go)
   Tests the concurrency and HTML-parsing logic.
2. Server Tests (tests/server_test.go)
   Uses httptest to validate endpoints like /, /data.
3. Database Tests (tests/database_test.go)
   Verifies database connectivity and table operations.

Run them all:
   go test ./...

For detailed test output:
   go test -v ./...

-------------------------------------------------------------------

Project Structure (Overview)

- cmd/webscraper/main.go  
  Entry point. Initializes the database and starts the server.  
- internal/server  
  Defines the HTTP server, routes (/scrape, /data), and handlers.  
- internal/scraper  
  Scrapes URLs concurrently, extracts <title> tags.  
- internal/database  
  Manages the PostgreSQL connection and schema setup.  
- internal/models  
  Contains the ScrapedData struct for storing parsed data.  
- tests/  
  Unit and integration tests.

-------------------------------------------------------------------

Documentation

All packages and functions include Go doc comments. You can view them in the command line:
   go doc ./...
Or simply browse the comments in each .go file.

