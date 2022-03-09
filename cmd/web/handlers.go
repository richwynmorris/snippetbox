package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the body response
func home(w http.ResponseWriter, r *http.Request) {
	// Check is the url path matches '/' exactly. If it doesnt,
	// use the http.NotFound() function to return a 404 responose to the
	// client. Return from the handler so that the function exits
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello form Snippetbox"))
}

func showSnippet(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Query().Get("id")
	id, err := strconv.Atoi(param)

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
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
		http.Error(w, "Method Not Allowed", statusCode)
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
