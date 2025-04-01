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
