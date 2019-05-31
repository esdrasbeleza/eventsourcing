package main

import (
	"database/sql"
	"time"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"
	"github.com/google/uuid"
)

func StoreEvent(db *sql.DB, event person.PersonEvent) error {
	var (
		sqlStatement = "INSERT INTO person_events (id, person_id, event_type, timestamp, data) VALUES ($1, $2, $3, $4, $5)"
		data         = event.JSON()
		_, err       = db.Exec(sqlStatement, uuid.New(), uuid.New(), event.Type(), time.Now(), data)
	)

	return err
}
