package main

import (
	"log"
	"net/http"
)

// Define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the body response

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello form Snippetbox"))
}

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux/router,
	// then register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	// Default path treats patter "/" as a catch all. All requests to the server will
	//handled by the home function
	mux.HandleFunc("/", home)

	log.Println("Starting server on port 4000")

	// ListenAndServe function starts a new web server. It takes two paramters, the TCP
	// network address and the router
	err := http.ListenAndServe(":4000", mux)
	// If the ListenAndServe function returns an error, you can log the error by passing it
	// to the fatal function and exiting.
	log.Fatal(err)
}
