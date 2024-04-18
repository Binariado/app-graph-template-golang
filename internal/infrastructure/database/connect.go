package database

import (
	"context"
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect(ctx context.Context) *sql.DB {

	dataSource := os.Getenv("POSTGRES_URL")
	if dataSource == "" {
		panic("missing environment variable POSTGRES_URL")
	}

	db, err := sql.Open("postgres", dataSource)

	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	return db
}

func Close() {
	if DB != nil {
		if err := DB.Close(); err != nil {
			log.Fatal("Error closing database: ", err)
		}
	}
}
