package storage

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func Test_MemoryStorage(t *testing.T) {
	memorySuite := newStorageTestSuite(NewMemoryStorage())
	suite.Run(t, memorySuite)
}
