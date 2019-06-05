package application

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/esdrasbeleza/eventsourcing/backend/storage"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CanCreateAPerson(t *testing.T) {
	db := DB()
	handler := Handler(db)
	httpServer := httptest.NewServer(handler)
	defer httpServer.Close()

	body, _ := json.Marshal(map[string]string{
		"Name":  "Esdras",
		"Email": "test@test.com",
	})

	response, _ := http.Post(httpServer.URL+"/person", "application/json", bytes.NewReader(body))

	var responseMap map[string]string
	responseBody, _ := ioutil.ReadAll(response.Body)
	json.Unmarshal(responseBody, &responseMap)

	assert.Equal(t, http.StatusCreated, response.StatusCode)
	assert.Equal(t, "Esdras", responseMap["Name"])
	assert.Equal(t, "test@test.com", responseMap["Email"])

	personId := uuid.MustParse(responseMap["Id"])

	storage := storage.NewDatabaseStorage(db)
	storedPerson, _ := storage.FetchPerson(personId)

	assert.Equal(t, "Esdras", storedPerson.Name)
	assert.Equal(t, "test@test.com", storedPerson.Email)
}
