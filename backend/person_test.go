package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreatePersonGivesUsAName(t *testing.T) {
	var (
		event  = ChangePersonName{Name: "Esdras"}
		person = GetPerson([]ChangePersonName{event})
	)

	assert.Equal(t, "Esdras", person.Name)
}
