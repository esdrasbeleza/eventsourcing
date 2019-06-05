package storage

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/esdrasbeleza/eventsourcing/eventsourcing/person"
	"github.com/google/uuid"
)

func ConnectToDB(connectionString string) *sql.DB {
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		panic(err)
	}

	return db
}

type DatabaseStorage struct {
	db *sql.DB
}

func NewDatabaseStorage(db *sql.DB) *DatabaseStorage {
	return &DatabaseStorage{db}
}

func (s *DatabaseStorage) StoreEvent(personId uuid.UUID, events ...person.PersonEvent) error {
	var (
		sqlStatement = "INSERT INTO person_events (id, person_id, event_type, timestamp, data) VALUES ($1, $2, $3, $4, $5)"
	)

	tx, err := s.db.Begin()

	if err != nil {
		return err
	}

	for _, event := range events {
		var (
			data   = event.JSON()
			_, err = tx.Exec(sqlStatement, uuid.New(), personId, event.Type(), time.Now(), data)
		)

		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (s *DatabaseStorage) FetchPerson(personId uuid.UUID) (*person.Person, error) {
	var (
		query     = "SELECT person_id,event_type,data FROM person_events WHERE person_id = $1 ORDER BY \"timestamp\""
		rows, err = s.db.Query(query, personId)
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	currentPerson := &person.Person{Id: personId}

	for rows.Next() {
		var (
			personId  uuid.UUID
			eventType string
			payload   []byte
		)

		if err := rows.Scan(&personId, &eventType, &payload); err != nil {
			return nil, err
		}

		event, err := person.EventFromName(eventType)

		if err != nil {
			return nil, err
		}

		if err := json.Unmarshal(payload, &event); err != nil {
			panic(err)
		}

		event.Apply(currentPerson)
	}

	return currentPerson, nil
}
