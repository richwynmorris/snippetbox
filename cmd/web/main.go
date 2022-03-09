package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	// Define a new command-line flar with the name 'addr'
	// The value of the flag will be stored in the addr variable at runtime

	addr := flag.String("addr", ":4000", "HTTP network address")
	// parses the command line flag. Needs to be called before using or
	// will use the defaul value
	flag.Parse()

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

	log.Printf("Starting server on port %s", *addr)

	// ListenAndServe function starts a new web server. It takes two paramters, the TCP
	// network address and the router
	err := http.ListenAndServe(*addr, mux)
	// If the ListenAndServe function returns an error, you can log the error by passing it
	// to the fatal function and exiting.
	log.Fatal(err)
}
