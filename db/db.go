package db

import (
	"database/sql"
)

func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./paste.db")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Migrate(db *sql.DB) error {
	sqlSmt := `
	CREATE TABLE IF NOT EXISTS paste (
		id TEXT PRIMARY KEY,
		content TEXT,
		language TEXT
	);
	`

	_, err := db.Exec(sqlSmt)
	if err != nil {
		return err
	}

	return nil
}
