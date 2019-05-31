package person

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleChangePersonNameEvent = ChangePersonName{Name: "Esdras"}

func Test_ChangePersonNameWorks(t *testing.T) {
	person := GetPerson([]PersonEvent{sampleChangePersonNameEvent})

	assert.Equal(t, "Esdras", person.Name)
}

func Test_ChangePersonName_JSON(t *testing.T) {
	expected, _ := json.Marshal(map[string]interface{}{
		"Name": "Esdras",
	})

	assert.JSONEq(t, string(expected), string(sampleChangePersonNameEvent.JSON()))
}
