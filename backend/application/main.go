package application

import "net/http"

func Start() {
	var (
		db      = DB()
		handler = Handler(db)
	)

	http.ListenAndServe(":8080", handler)
}
