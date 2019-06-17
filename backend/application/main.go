package application

import (
	"log"
	"net/http"
	"os"
)

func Start() {
	var (
		db      = DB()
		handler = Handler(db)
		logger  = log.New(os.Stdout, "http: ", log.LstdFlags)
	)

	logger.Println("Server is starting...")

	server := http.Server{
		Addr:     ":8081",
		ErrorLog: logger,
		Handler:  handler,
	}

	log.Fatal(server.ListenAndServe())
}
