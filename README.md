# Go Web Scraper

A simple web scraper built with Go to demonstrate fundamental Go features:
- Building a web server
- Concurrency with goroutines
- Database connectivity
- Basic testing

## Getting Started

1. **Clone the repo**:
   ```bash
   git clone https://github.com/Olt-Kondirolli91/go-web-scraper.git
   cd go-web-scraper
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set up the Database** (PostgreSQL in this example):
   - Update the database URL in `cmd/webscraper/main.go`.
   - Make sure your PostgreSQL instance is running.

4. **Run the Server**:
   ```bash
   go run ./cmd/webscraper
   ```
   - Opens on [http://localhost:8080](http://localhost:8080).

## API Endpoints

- **GET /**  
  Returns a welcome message.

- **GET /scrape**  
  Scrapes sample URLs and stores their titles in the database.

- **GET /data**  
  Returns the scraped data in JSON format.

## Testing

Run all tests:
```bash
go test ./...
```
If you want more detailed output, use:
```bash
go test -v ./...
```

## Project Structure

- `cmd/webscraper/main.go` – Entry point
- `internal/server` – Server and route setup
- `internal/scraper` – Concurrency and HTML parsing logic
- `internal/database` – Database connection and setup
- `internal/models` – Data models
- `tests` – Unit and integration tests

