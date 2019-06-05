package storage

import (
	"testing"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CanStoreEvent(t *testing.T) {
	var (
		dbStorage = NewDatabaseStorage(db)
		personId         = uuid.New()
		changePersonName = person.ChangePersonName{Name: "Esdras"}
		err              = dbStorage.StoreEvent(personId, changePersonName)
	)

	assert.Nil(t, err)
}

func Test_Can_FetchEvents(t *testing.T) {
	var (
		dbStorage = NewDatabaseStorage(db)

		personId = uuid.New()
		event1   = person.ChangePersonName{Name: "Esdras"}
		event2   = person.AddEmail{Email: "test@test.com"}
		event3   = person.AddAddress{Name: "Home", Address: "My address"}
		event4   = person.AddAddress{Name: "Work", Address: "My work address"}
	)

	dbStorage.StoreEvent(personId, event1)
	dbStorage.StoreEvent(personId, event2)
	dbStorage.StoreEvent(personId, event3)
	dbStorage.StoreEvent(personId, event4)

	person, err := dbStorage.FetchPerson(personId)

	assert.Nil(t, err)

	assert.Equal(t, "Esdras", person.Name)
	assert.Equal(t, "test@test.com", person.Email)
	assert.Equal(t, "My address", person.Address["Home"])
	assert.Equal(t, "My work address", person.Address["Work"])
}
