package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

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
