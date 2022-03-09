package main

import (
	"log"
	"net/http"
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
	w.Write([]byte("Display a specific snippet..."))
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

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux/router,
	// then register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	// Default path treats patter "/" as a catch all. All requests to the server will
	//handled by the home function
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on port 4000")

	// ListenAndServe function starts a new web server. It takes two paramters, the TCP
	// network address and the router
	err := http.ListenAndServe(":4000", mux)
	// If the ListenAndServe function returns an error, you can log the error by passing it
	// to the fatal function and exiting.
	log.Fatal(err)
}
