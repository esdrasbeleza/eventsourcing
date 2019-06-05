package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"
	"github.com/esdrasbeleza/eventsourcing/eventsourcing/storage"

	"github.com/google/uuid"
)

type PersonController struct {
	Storage storage.Person
}

func (c *PersonController) CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var input struct {
		Name  string
		Email string
	}

	if err := json.Unmarshal(requestBody, &input); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var (
		uuid       = uuid.New()
		changeName = person.ChangePersonName{Name: input.Name}
		addEmail   = person.AddEmail{Email: input.Email}
	)

	if err := c.Storage.StoreEvent(uuid, changeName, addEmail); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := person.GetPerson([]person.PersonEvent{changeName, addEmail})
	response.Id = uuid

	responseBody, _ := json.Marshal(response)

	w.WriteHeader(http.StatusCreated)
	w.Write(responseBody)
}
