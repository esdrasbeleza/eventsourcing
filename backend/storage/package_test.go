package storage

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	defer afterTests()
	beforeTests()
	code := m.Run()
	os.Exit(code)
}

func beforeTests() {
	db = ConnectToDB("postgres://postgres:docker@localhost/postgres?sslmode=disable")
	checkConnection()
}

func afterTests() {
	if db != nil {
		db.Close()
	}
}

func checkConnection() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
