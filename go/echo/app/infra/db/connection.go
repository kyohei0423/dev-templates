package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq" // postgres driver
)

var db *sql.DB

func NewDBconnection() *sql.DB {
	if db != nil {
		return db
	}

	connStr := os.Getenv("POSTGRES_URL")
	if connStr == "" {
		connStr = "postgres://root:password@db:5432/sample?sslmode=disable" // TODO: Fix DB name
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic("Failed to connect Database")
	}

	connErr := db.Ping()
	if connErr != nil {
		panic("Database is not alive")
	}

	return db
}
