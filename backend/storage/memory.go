package storage

import (
	"errors"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"

	"github.com/google/uuid"
)

type MemoryStorage struct {
	storage map[uuid.UUID][]person.PersonEvent
}

func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		storage: make(map[uuid.UUID][]person.PersonEvent),
	}
}

func (m *MemoryStorage) StoreEvent(personId uuid.UUID, event person.PersonEvent) error {
	if _, exists := m.storage[personId]; !exists {
		m.storage[personId] = []person.PersonEvent{}
	}

	m.storage[personId] = append(m.storage[personId], event)

	return nil
}

func (m *MemoryStorage) FetchPerson(personId uuid.UUID) (*person.Person, error) {
	events, exists := m.storage[personId]

	if !exists {
		return nil, errors.New("Not found")
	}

	person := person.GetPerson(events)

	return person, nil
}
