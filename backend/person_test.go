package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePersonGivesUsAName(t *testing.T) {
	var (
		event  = CreatePerson{Name: "Esdras"}
		person = GetPerson([]CreatePerson{event})
	)

	assert.Equal(t, "Esdras", person.Name)
}
