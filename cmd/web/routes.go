package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	//Create a middleware chain containing our standard middlewares
	standardMiddleware := alice.New(app.recoverPanic, app.logRequest, secureHeaders)

	// Use the http.NewServeMux() function to initialize a new servemux/router,
	// then register the home function as the handler for the "/" URL pattern
	mux := pat.New()
	// Default path treats patter "/" as a catch all. All requests to the server will
	//handled by the home function
	mux.Get("/", http.HandlerFunc(app.home))
	mux.Get("/snippet/create", http.HandlerFunc(app.createSnippetForm))
	mux.Post("/snippet/create", http.HandlerFunc(app.createSnippet))
	mux.Get("/snippet/:id", http.HandlerFunc(app.showSnippet))

	// Create a file server which serves files out of the static directory
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))

	return standardMiddleware.Then(mux)
}
