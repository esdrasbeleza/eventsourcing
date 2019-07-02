package person

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleRemoveAddressEvent = RemoveAddress{Name: "Home"}

func Test_RemoveAddress(t *testing.T) {
	person := GetPerson([]PersonEvent{sampleAddAddressEvent, sampleRemoveAddressEvent})

	assert.Empty(t, person.Address["Home"])
}

func Test_RemoveAddress_JSON(t *testing.T) {
	expected, _ := json.Marshal(map[string]interface{}{
		"Name": "Home",
	})

	assert.JSONEq(t, string(expected), string(sampleRemoveAddressEvent.JSON()))
}
