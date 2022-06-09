package main

import (
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"

	"richwynmorris.co.uk/internal/models"
)

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	snippets      *models.SnippetModel
	templateCache map[string]*template.Template
	formDecoder   *form.Decoder
}

func main() {
	// Initialize cli arguments for flags.
	addr := flag.String("addr", ":4000", "HTTP Network Address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()

	// Initialize loggers.
	infoLog := log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Open database connection.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// Defer closing of database connection for graceful shutdown.
	defer db.Close()

	// Initialize snippetModel with open database connection.
	snippetModel := &models.SnippetModel{DB: db}

	// Initialize a templateCache to be used for html rendering.
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize decoder instance...
	formDecoder := form.NewDecoder()

	// Inject looger and database dependencies into application struct
	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		snippets:      snippetModel,
		templateCache: templateCache,
		formDecoder:   formDecoder,
	}

	// Initialize server with handler to access the routes available on the app
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

// openDB receives the data source name, opens a database connection, ping's database to check connection is alive and
// return open database connection.
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
