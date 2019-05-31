package person

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_EventFromName(t *testing.T) {
	var (
		eventName    = "ChangePersonName"
		event, _     = EventFromName(eventName)
		eventPayload = `{"Name": "Esdras"}`
	)

	json.Unmarshal([]byte(eventPayload), &event)

	person := &Person{}
	event.Apply(person)

	assert.Equal(t, "Esdras", person.Name)
}
