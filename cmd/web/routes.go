package main

import "net/http"

func (app *application) routes() http.Handler {
	// Multiplexer
	mux := http.NewServeMux()

	// STATIC FILES HANDLER
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	// HANDLERS
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	return app.recoverPanic(app.logRequest(secureHeaders(mux)))
}
