package person

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
