package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddAddressWorks(t *testing.T) {
	var (
		event  = AddAddress{Name: "Home", Address: "Some address"}
		person = GetPerson([]PersonEvent{event})
	)

	assert.Equal(t, "Some address", person.Address["Home"])
}
