package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AddEmailWorks(t *testing.T) {
	var (
		event  = AddEmail{Email: "test@test.com"}
		person = GetPerson([]PersonEvent{event})
	)

	assert.Equal(t, "test@test.com", person.Email)
}
