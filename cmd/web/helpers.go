package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helps write an error message and stack trace to the errorLog
// Then sends a generic 500 response to the use

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// clientError send specific status code and description for 400 responses.
func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

// Not found helper - wraps around clientError but passes down 404 status to be returned
func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
