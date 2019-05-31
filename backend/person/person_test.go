package person

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPerson_HasAnId(t *testing.T) {
	person := NewPerson()

	assert.NotEmpty(t, person.Id.String())
}
