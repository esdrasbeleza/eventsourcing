package application

import (
	"database/sql"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/storage"
)

func DB() *sql.DB {
	return storage.ConnectToDB("postgres://postgres:docker@localhost/postgres?sslmode=disable")
}
