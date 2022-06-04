package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet/view", app.snippetView)
	mux.HandleFunc("/snippet/create", app.snippetCreate)

	// Use Alice package to create middleware chain and add readability.
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Return the standardMiddleware chain followed by servemux.
	return standardMiddleware.Then(mux)
}
