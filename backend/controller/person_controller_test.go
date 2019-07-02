package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/esdrasbeleza/eventsourcing/backend/person"
	"github.com/esdrasbeleza/eventsourcing/backend/storage"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_ItCanCreateAPerson(t *testing.T) {
	var (
		storage    = storage.NewMemoryStorage()
		controller = &PersonController{storage}

		requestBody, _ = json.Marshal(map[string]string{
			"Name":  "Esdras",
			"Email": "test@test.com",
		})

		request, _ = http.NewRequest(http.MethodPost, "/person", bytes.NewReader(requestBody))

		recorder = httptest.NewRecorder()
	)

	controller.CreatePerson(recorder, request)

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)

	var responseJSON map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &responseJSON)

	assert.NotPanics(t, func() { uuid.MustParse(responseJSON["Id"]) })
	assert.Equal(t, "Esdras", responseJSON["Name"])
	assert.Equal(t, "test@test.com", responseJSON["Email"])
}

func Test_ItCanReadAPerson(t *testing.T) {
	var (
		uuid       = uuid.New()
		storage    = storage.NewMemoryStorage()
		controller = &PersonController{storage}

		events = []person.PersonEvent{
			person.ChangePersonName{Name: "Esdras"},
			person.AddEmail{Email: "test@test.com"},
		}
	)

	storage.StoreEvent(uuid, events...)

	var (
		request, _ = http.NewRequest(http.MethodGet, "/person/"+uuid.String(), nil)
		recorder   = httptest.NewRecorder()
	)

	request = mux.SetURLVars(request, map[string]string{"id": uuid.String()})

	controller.GetPerson(recorder, request)

	var responseJSON map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusOK, recorder.Result().StatusCode)
	assert.Equal(t, "Esdras", responseJSON["Name"])
	assert.Equal(t, "test@test.com", responseJSON["Email"])
}

func Test_AddAddress_ReturnsExpectedAddress(t *testing.T) {
	var (
		uuid       = uuid.New()
		storage    = storage.NewMemoryStorage()
		controller = &PersonController{storage}

		events = []person.PersonEvent{
			person.ChangePersonName{Name: "Esdras"},
			person.AddEmail{Email: "test@test.com"},
		}
	)

	storage.StoreEvent(uuid, events...)

	requestBody, _ := json.Marshal(map[string]string{
		"Name":    "Home",
		"Address": "Address",
	})

	var (
		request, _ = http.NewRequest(http.MethodPost, "/person/"+uuid.String()+"/address", bytes.NewReader(requestBody))
		recorder   = httptest.NewRecorder()
	)

	request = mux.SetURLVars(request, map[string]string{"id": uuid.String()})

	controller.AddAddress(recorder, request)

	var responseJSON map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &responseJSON)

	assert.Equal(t, http.StatusCreated, recorder.Result().StatusCode)
	assert.Equal(t, "Home", responseJSON["Name"])
	assert.Equal(t, "Address", responseJSON["Address"])

	updatedPerson, _ := storage.FetchPerson(uuid)
	assert.Equal(t, "Address", updatedPerson.Address["Home"])
}
