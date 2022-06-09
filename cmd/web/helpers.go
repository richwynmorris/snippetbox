package main

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/go-playground/form/v4"
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

	// Initialize a new buffer which uses the same interface as http.ResponseWriter
	buf := new(bytes.Buffer)

	// Execute the template set and pass in the data to be dynamically rendered

	// The executed template should be handed to the buffer instead of http.ResponseWriter
	// to check for any errors when executing the template.
	err := ts.ExecuteTemplate(buf, "base", data)
	if err != nil {
		app.serverError(w, err)
	}

	// If no errors have occurred, we're safe to write out the provided http status code
	w.WriteHeader(status)

	// The contents of the buffer should now be written to the http.ResponseWriter
	buf.WriteTo(w)
}

func (app *application) newTemplateData(r *http.Request) *templateData {
	return &templateData{
		CurrentYear: time.Now().Year(),
	}
}

func (app *application) decodePostForm(r *http.Request, dest any) error {
	err := r.ParseForm()
	if err != nil {
		return err
	}

	err = app.formDecoder.Decode(dest, r.PostForm)
	if err != nil {
		var invalidDecoderError *form.InvalidDecoderError
		if errors.As(err, &invalidDecoderError) {
			panic(err)
		}

		return err
	}

	return nil
}
