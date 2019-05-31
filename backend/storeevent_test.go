package main

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func Test_CanStoreEvent(t *testing.T) {
	var (
		changePersonName = person.ChangePersonName{Name: "Esdras"}
		sqlStatement     = "INSERT INTO person_events (id, person_id, event_type, timestamp, data) VALUES ($1, $2, $3, $4, $5)"
		data, _          = json.Marshal(map[string]interface{}{"name": changePersonName.Name})
		_, err           = db.Exec(sqlStatement, uuid.New(), uuid.New(), "change_person_name", time.Now(), data)
	)

	assert.Nil(t, err)
}
