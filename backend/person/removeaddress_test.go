package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RemoveAddress(t *testing.T) {
	var (
		event1 = AddAddress{Name: "Home", Address: "Some address"}
		event2 = RemoveAddress{Name: "Home"}
		person = GetPerson([]PersonEvent{event1, event2})
	)

	assert.Nil(t, person.Address["Home"])
}
