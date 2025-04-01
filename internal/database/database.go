package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

// ConnectPostgres connects to a PostgreSQL database using the given DSN. 
// It also ensures the scraped_data table exists. 
// Returns the open *sql.DB or an error if the connection fails.
func ConnectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Create table if it doesn't exist
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS scraped_data (
			id SERIAL PRIMARY KEY,
			url TEXT NOT NULL,
			title TEXT
		);
	`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

