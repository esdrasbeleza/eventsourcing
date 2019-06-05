package storage

import (
	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type storageTestSuite struct {
	suite.Suite
	storage Person
}

func newStorageTestSuite(storage Person) *storageTestSuite {
	return &storageTestSuite{
		storage: storage,
	}
}

func (s storageTestSuite) Test_CanStoreEvent() {
	var (
		personId         = uuid.New()
		changePersonName = person.ChangePersonName{Name: "Esdras"}
		err              = s.storage.StoreEvent(personId, changePersonName)
	)

	s.Nil(err)
}

func (s storageTestSuite) Test_CanFetchEvents() {
	var (
		personId = uuid.New()
		event1   = person.ChangePersonName{Name: "Esdras"}
		event2   = person.AddEmail{Email: "test@test.com"}
		event3   = person.AddAddress{Name: "Home", Address: "My address"}
		event4   = person.AddAddress{Name: "Work", Address: "My work address"}
	)

	s.storage.StoreEvent(personId, event1, event2, event3, event4)

	person, err := s.storage.FetchPerson(personId)

	s.Nil(err)

	s.Equal("Esdras", person.Name)
	s.Equal("test@test.com", person.Email)
	s.Equal("My address", person.Address["Home"])
	s.Equal("My work address", person.Address["Work"])
}
