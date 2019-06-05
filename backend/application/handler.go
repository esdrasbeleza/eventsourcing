package application

import (
	"database/sql"
	"net/http"

	"github.com/esdrasbeleza/eventsourcing/backend/controller"
	"github.com/esdrasbeleza/eventsourcing/backend/storage"
	"github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

func Handler(db *sql.DB) http.Handler {
	var (
		storage    = storage.NewDatabaseStorage(db)
		controller = controller.PersonController{storage}
		router     = mux.NewRouter()
	)

	router.HandleFunc("/person", controller.CreatePerson).Methods(http.MethodPost)

	return router
}
