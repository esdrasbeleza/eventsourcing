package storage

import (
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/suite"
)

func Test_DatabaseStorage(t *testing.T) {
	db := ConnectToDB("postgres://postgres:docker@localhost/postgres?sslmode=disable")
	defer db.Close()

	memorySuite := newStorageTestSuite(NewDatabaseStorage(db))
	suite.Run(t, memorySuite)
}
