package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"
	"github.com/google/uuid"
)

type PersonStorage interface {
	StoreEvent(personId uuid.UUID, event person.PersonEvent) error
	FetchPerson(personId uuid.UUID) (*person.Person, error)
}

type PersonController struct {
	storage PersonStorage
}

func (c *PersonController) CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var input person.ChangePersonName

	if err := json.Unmarshal(requestBody, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	uuid := uuid.New()

	if err := c.storage.StoreEvent(uuid, input); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := person.GetPerson([]person.PersonEvent{input})
	response.Id = uuid

	responseBody, _ := json.Marshal(response)

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
