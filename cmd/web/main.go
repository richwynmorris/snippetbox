package main

import (
	"log"
	"net/http"
)

func main() {
	// Use the http.NewServeMux() function to initialize a new servemux/router,
	// then register the home function as the handler for the "/" URL pattern
	mux := http.NewServeMux()
	// Default path treats patter "/" as a catch all. All requests to the server will
	//handled by the home function
	mux.HandleFunc("/", home)

	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the static directory
	fileServer := http.FileServer(http.Dir("./ui/static"))

	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on port 4000")

	// ListenAndServe function starts a new web server. It takes two paramters, the TCP
	// network address and the router
	err := http.ListenAndServe(":4000", mux)
	// If the ListenAndServe function returns an error, you can log the error by passing it
	// to the fatal function and exiting.
	log.Fatal(err)
}
