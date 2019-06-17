package application

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/esdrasbeleza/eventsourcing/backend/storage"
)

func DB() *sql.DB {
	dbHost := os.Getenv("DATABASE_HOST")

	if dbHost == "" {
		dbHost = "localhost"
	}

	log.Printf("Database is %s:5432\n", dbHost)

	dbURL := fmt.Sprintf("postgres://postgres:docker@%s/postgres?sslmode=disable", dbHost)

	return storage.ConnectToDB(dbURL)
}
