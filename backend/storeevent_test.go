package main

import (
	"testing"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CanStoreEvent(t *testing.T) {
	var (
		personId         = uuid.New()
		changePersonName = person.ChangePersonName{Name: "Esdras"}
		err              = StoreEvent(db, personId, changePersonName)
	)

	assert.Nil(t, err)
}

func Test_Can_FetchEvents(t *testing.T) {
	var (
		personId = uuid.New()
		event1   = person.ChangePersonName{Name: "Esdras"}
		event2   = person.AddEmail{Email: "test@test.com"}
		event3   = person.AddAddress{Name: "Home", Address: "My address"}
		event4   = person.AddAddress{Name: "Work", Address: "My work address"}
	)

	StoreEvent(db, personId, event1)
	StoreEvent(db, personId, event2)
	StoreEvent(db, personId, event3)
	StoreEvent(db, personId, event4)

	person, err := FetchPerson(db, personId)

	assert.Nil(t, err)

	assert.Equal(t, "Esdras", person.Name)
	assert.Equal(t, "test@test.com", person.Email)
	assert.Equal(t, "My address", person.Address["Home"])
	assert.Equal(t, "My work address", person.Address["Work"])
}
