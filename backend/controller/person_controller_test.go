package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"

	"github.com/esdrasbeleza/eventsourcing/backend/person"
	"github.com/esdrasbeleza/eventsourcing/backend/storage"
)

func TestPersonController(t *testing.T) {
	suite.Run(t, new(PersonControlerSuite))
}

type PersonControlerSuite struct {
	suite.Suite
}

func (s *PersonControlerSuite) Test_ItCanCreateAPerson() {
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

	s.Equal(http.StatusCreated, recorder.Result().StatusCode)

	var responseJSON map[string]string
	json.Unmarshal(recorder.Body.Bytes(), &responseJSON)

	s.NotPanics(func() { uuid.MustParse(responseJSON["Id"]) })
	s.Equal("Esdras", responseJSON["Name"])
	s.Equal("test@test.com", responseJSON["Email"])
}

func (s *PersonControlerSuite) Test_ItCanReadAPerson() {
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

	s.Equal(http.StatusOK, recorder.Result().StatusCode)
	s.Equal("Esdras", responseJSON["Name"])
	s.Equal("test@test.com", responseJSON["Email"])
}

func (s *PersonControlerSuite) Test_AddAddress_ReturnsExpectedAddress() {
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

	s.Equal(http.StatusCreated, recorder.Result().StatusCode)
	s.Equal("Home", responseJSON["Name"])
	s.Equal("Address", responseJSON["Address"])

	updatedPerson, _ := storage.FetchPerson(uuid)
	s.Equal("Address", updatedPerson.Address["Home"])
}
