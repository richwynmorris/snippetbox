package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	// Define a new command-line flar with the name 'addr'
	// The value of the flag will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")
	// parses the command line flag. Needs to be called before using or
	// will use the defaul value
	flag.Parse()
	// Use log.New() to create a logger. Takes 3 params - destination, prefix and flags
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// For error logs, use stderr as dest and log.Lshortfile
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//-------------------------------------------------------------------------------------

	// Initializes new application struct - has errorLog and InfoLog dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	// Initiazlise new http.Server struct - add address, error log and handler dependencies
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	//-------------------------------------------------------------------------------------

	infoLog.Printf("Starting server on port %s", *addr)

	// ListenAndServe function starts a new web server. Called from new http.Server struct
	err := server.ListenAndServe()

	// If the ListenAndServe function returns an error, you can log the error by passing it
	// to the fatal function and exiting.
	errorLog.Fatal(err)
}
