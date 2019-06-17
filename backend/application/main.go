package application

import (
	"log"
	"net/http"
)

func Start() {
	var (
		db      = DB()
		handler = Handler(db)
	)

	log.Fatal(http.ListenAndServe(":8080", handler))
}
