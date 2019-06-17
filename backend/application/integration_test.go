package application

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esdrasbeleza/eventsourcing/backend/storage"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

func TestPersonController(t *testing.T) {
	suite.Run(t, new(PersonControlerSuite))
}

type PersonControlerSuite struct {
	suite.Suite
	db         *sql.DB
	handler    http.Handler
	storage    *storage.DatabaseStorage
	httpServer *httptest.Server
}

func (s *PersonControlerSuite) SetupSuite() {
	s.db = DB()
	s.handler = Handler(s.db)
	s.storage = storage.NewDatabaseStorage(s.db)
	s.httpServer = httptest.NewServer(s.handler)
}

func (s *PersonControlerSuite) TearDownSuite() {
	if s.httpServer != nil {
		s.httpServer.Close()
	}
}

func (s *PersonControlerSuite) createPerson(body []byte) (*http.Response, map[string]string) {
	response, _ := http.Post(s.httpServer.URL+"/person", "application/json", bytes.NewReader(body))

	var responseMap map[string]string
	responseBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseBody, &responseMap)

	return response, responseMap
}

func (s *PersonControlerSuite) Test_CanCreateAPerson() {
	body, _ := json.Marshal(map[string]string{
		"Name":  "Esdras",
		"Email": "test@test.com",
	})

	response, responseMap := s.createPerson(body)

	s.Equal(http.StatusCreated, response.StatusCode)
	s.Equal("Esdras", responseMap["Name"])
	s.Equal("test@test.com", responseMap["Email"])

	personId := uuid.MustParse(responseMap["Id"])

	storedPerson, _ := s.storage.FetchPerson(personId)

	s.Equal("Esdras", storedPerson.Name)
	s.Equal("test@test.com", storedPerson.Email)
}

func (s *PersonControlerSuite) Test_CanReadAPerson() {
	body, _ := json.Marshal(map[string]string{
		"Name":  "Beleza",
		"Email": "test2@test.com",
	})

	_, createdPerson := s.createPerson(body)

	getUserResponse, _ := http.Get(s.httpServer.URL + "/person/" + createdPerson["Id"])

	var responseMap map[string]string
	responseBody, _ := ioutil.ReadAll(getUserResponse.Body)
	json.Unmarshal(responseBody, &responseMap)

	s.Equal("Beleza", responseMap["Name"])
	s.Equal("test2@test.com", responseMap["Email"])
}
