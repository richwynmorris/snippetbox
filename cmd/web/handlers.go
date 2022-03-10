package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
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
	err = ts.Execute(w, nil)
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

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	const statusCode int = 405
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
	w.Write([]byte("Create a new snippet..."))
}
