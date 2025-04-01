package database

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func ConnectPostgres(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS scraped_data (
		id SERIAL PRIMARY KEY,
		url TEXT NOT NULL,
		title TEXT
	)`)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func SeedData(db *sql.DB) error {
	_, err := db.Exec(`
		INSERT INTO scraped_data (url, title)
		VALUES 
			('https://mysite.com', 'My Website'),
			('https://anothersite.io', 'Another Site'),
			('https://example.org', 'Example Org')
		ON CONFLICT DO NOTHING; -- If you have a unique constraint and don't want duplicates
	`)
	return err
}
