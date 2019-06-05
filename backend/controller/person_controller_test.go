package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func router() *mux.Router {
	router := mux.NewRouter()
	return router
}

type memoryStorage struct {
	storage map[uuid.UUID][]person.PersonEvent
}

func newMemoryStorage() memoryStorage {
	return memoryStorage{
		storage: make(map[uuid.UUID][]person.PersonEvent),
	}
}

func (m memoryStorage) StoreEvent(personId uuid.UUID, event person.PersonEvent) error {
	if _, exists := m.storage[personId]; !exists {
		m.storage[personId] = []person.PersonEvent{}
	}

	m.storage[personId] = append(m.storage[personId], event)

	return nil
}

func (m memoryStorage) FetchPerson(personId uuid.UUID) (*person.Person, error) {
	events, exists := m.storage[personId]

	if !exists {
		return nil, errors.New("Not found")
	}

	person := person.GetPerson(events)

	return person, nil
}

func Test_ItCanCreateAPerson(t *testing.T) {
	var (
		storage    = newMemoryStorage()
		controller = &PersonController{storage}

		requestBody, _ = json.Marshal(map[string]string{"Name": "Esdras"})
		request, _     = http.NewRequest(http.MethodPost, "/person", bytes.NewReader(requestBody))

		recorder = httptest.NewRecorder()
	)

	controller.CreatePerson(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	var responseJSON map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &responseJSON)

	assert.Equal(t, "Esdras", responseJSON["Name"])
	assert.NotPanics(t, func() { uuid.MustParse(responseJSON["Id"]) })
}
