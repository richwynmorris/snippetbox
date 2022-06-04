package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, status int, page string, data *templateData) {
	// check if the page is in the template cache
	ts, ok := app.templateCache[page]

	// if ok returns false, generate an error and pass it to the serverError func
	if !ok {
		err := fmt.Errorf("the template %s does not exists", page)
		app.serverError(w, err)
	}

	// Write out the provided http status code
	w.WriteHeader(status)

	// Execute the template set and pass in the data to be dynamically rendered
	err := ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		app.serverError(w, err)
	}
}
