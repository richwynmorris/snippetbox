package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"richwynmorris.co.uk/snippetbox/pkg/models/mysql"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
}

func main() {
	// Define a new command-line flar with the name 'addr'
	// The value of the flag will be stored in the addr variable at runtime
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Define a new command-line flad for the MRSQL DSN string
	dsn := flag.String("dsn", "web:jjrchorus21@/snippetbox?parseTime=true", "MYSQL data source name")

	// parses the command line flag. Needs to be called before using or
	// will use the defaul value
	flag.Parse()

	// Use log.New() to create a logger. Takes 3 params - destination, prefix and flags
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// For error logs, use stderr as dest and log.Lshortfile
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//-------------------------------------------------------------------------------------

	// Pass in DSN from the command line flag to the connection pool via `openDB()`
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// call defer db.Close() so that the connection pool is closed befoer the main function
	// exits
	defer db.Close()

	// Initializes new application struct - has errorLog and InfoLog dependencies
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
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
	err = server.ListenAndServe()

	// If the ListenAndServe function returns an error, you can log the error by passing it
	// to the fatal function and exiting.
	errorLog.Fatal(err)
}

// openDB function wraps sql.Open and returns the connection ppol for
// a given DSN
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
