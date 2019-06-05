package storage

import (
	"database/sql"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var db *sql.DB

func TestMain(m *testing.M) {
	defer testShutdown()
	testSetup()
	code := m.Run()
	os.Exit(code)
}

func testSetup() {
	connectToTestDB()
	checkConnection()
}

func testShutdown() {
	if db != nil {
		db.Close()
	}
}

func connectToTestDB() {
	connStr := "postgres://postgres:docker@localhost/postgres?sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}
}

func checkConnection() {
	if err := db.Ping(); err != nil {
		panic(err)
	}
}
