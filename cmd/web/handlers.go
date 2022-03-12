package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"

	"richwynmorris.co.uk/snippetbox/pkg/models"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the body response
func (app *application) home(w http.ResponseWriter, r *http.Request) {
	// Check is the url path matches '/' exactly. If it doesnt,
	// use the http.NotFound() function to return a 404 responose to the
	// client. Return from the handler so that the function exits
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	snippets, err := app.snippets.Latest()
	if err != nil {
		app.serverError(w, err)
	}

	data := &templateData{Snippets: snippets}

	// initialize a slice containing th paths to the two templating files
	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	// ParseFiles reads the template file into a template set
	ts, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	// Execute method on template set write the template context as
	// the response body. The second argument takes in any dynamic data
	// that is relevant to the
	err = ts.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	id, err := strconv.Atoi(param)

	if err != nil || id < 1 {
		app.notFound(w)
		return
	}

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.notFound(w)
		} else {
			app.serverError(w, err)
		}
		return
	}

	data := &templateData{
		Snippet: snippet,
	}

	files := []string{
		"./ui/html/show.page.tmpl",
		"./ui/html/base.layout.tmpl",
		"./ui/html/footer.partial.tmpl",
	}

	templates, err := template.ParseFiles(files...)
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = templates.Execute(w, data)
	if err != nil {
		app.serverError(w, err)
	}

	fmt.Fprintf(w, "%v", snippet)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	// Use r.Method to check whether the request is using a POST or not.
	if r.Method != "POST" {
		// If the method is not POST, send a response with the status 405
		// (METHOD not accepted) in the response body. Added an 'Allow: POST' header to the
		// response. The first param is the header name and the second is the header value
		// This must be called before either WriteHeader or Write methods
		w.Header().Set("Allow", "POST")
		// Use http.Error() function to send a 405 status code and "Method Not Allowed"
		// string as the response body.
		app.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	// Dummy data:
	title := "0 snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	// Pass data to SnippetModel.Insert() - returns the ID of newly created record
	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.serverError(w, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet?id=%d", id), http.StatusSeeOther)
}
