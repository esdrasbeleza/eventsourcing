package storage

import (
	"github.com/esdrasbeleza/eventsourcing/backend/person"
	"github.com/google/uuid"
)

type Person interface {
	StoreEvent(personId uuid.UUID, events ...person.PersonEvent) error
	FetchPerson(personId uuid.UUID) (*person.Person, error)
}
