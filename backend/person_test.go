package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ChangePersonNameWorks(t *testing.T) {
	var (
		event  = ChangePersonName{Name: "Esdras"}
		person = GetPerson([]PersonEvent{event})
	)

	assert.Equal(t, "Esdras", person.Name)
}

func Test_AddEmailWorks(t *testing.T) {
	var (
		event  = AddEmail{Email: "test@test.com"}
		person = GetPerson([]PersonEvent{event})
	)

	assert.Equal(t, "test@test.com", person.Email)
}

func Test_AddAddressWorks(t *testing.T) {
	var (
		event  = AddAddress{Name: "Home", Address: "Some address"}
		person = GetPerson([]PersonEvent{event})
	)

	assert.Equal(t, "Some address", person.Address["Home"])
}

func Test_RemoveAddress(t *testing.T) {
	var (
		event1 = AddAddress{Name: "Home", Address: "Some address"}
		event2 = RemoveAddress{Name: "Home"}
		person = GetPerson([]PersonEvent{event1, event2})
	)

	assert.Nil(t, person.Address["Home"])
}