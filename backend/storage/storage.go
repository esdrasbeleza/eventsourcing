package storage

import (
	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"
	"github.com/google/uuid"
)

type Person interface {
	StoreEvent(personId uuid.UUID, events ...person.PersonEvent) error
	FetchPerson(personId uuid.UUID) (*person.Person, error)
}