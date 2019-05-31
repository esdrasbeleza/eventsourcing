package person

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleAddAddressEvent = AddAddress{
	Name:    "Home",
	Address: "Some address",
}

func Test_AddAddressWorks(t *testing.T) {
	person := GetPerson([]PersonEvent{sampleAddAddressEvent})

	assert.Equal(t, "Some address", person.Address["Home"])
}

func Test_AddAddress_JSON(t *testing.T) {
	expected, _ := json.Marshal(map[string]interface{}{
		"Name":    "Home",
		"Address": "Some address",
	})

	assert.JSONEq(t, string(expected), string(sampleAddAddressEvent.JSON()))
}
