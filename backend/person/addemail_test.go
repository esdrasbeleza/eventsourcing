package person

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleAddEmailEvent = AddEmail{Email: "test@test.com"}

func Test_AddEmailWorks(t *testing.T) {
	person := GetPerson([]PersonEvent{sampleAddEmailEvent})

	assert.Equal(t, "test@test.com", person.Email)
}

func Test_AddEmail_JSON(t *testing.T) {
	expected, _ := json.Marshal(map[string]interface{}{
		"Email": "test@test.com",
	})

	assert.JSONEq(t, string(expected), string(sampleAddEmailEvent.JSON()))
}
