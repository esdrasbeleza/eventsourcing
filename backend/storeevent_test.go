package main

import (
	"testing"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"

	"github.com/stretchr/testify/assert"
)

func Test_CanStoreEvent(t *testing.T) {
	var (
		changePersonName = person.ChangePersonName{Name: "Esdras"}
		err              = StoreEvent(db, changePersonName)
	)

	assert.Nil(t, err)
}
