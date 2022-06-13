package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"

	"richwynmorris.co.uk/internal/models"
)

type application struct {
	errorLog       *log.Logger
	formDecoder    *form.Decoder
	infoLog        *log.Logger
	sessionManager *scs.SessionManager
	snippets       *models.SnippetModel
	templateCache  map[string]*template.Template
	users          *models.UserModel
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
	// Initialize UserModel with open data database connection
	userModel := &models.UserModel{DB: db}

	// Initialize a templateCache to be used for html rendering.
	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	// Initialize decoder instance...
	formDecoder := form.NewDecoder()

	// Initialize session manager.
	sessionManager := scs.New()
	// Configure sessions manager to use MySQL db.
	sessionManager.Store = mysqlstore.New(db)
	// Set session's expire date.
	sessionManager.Lifetime = 12 * time.Hour
	// Set HTTPS connection only for session data.
	sessionManager.Cookie.Secure = true

	// Application struct containing the app's dependencies.
	app := &application{
		errorLog:       errorLog,
		formDecoder:    formDecoder,
		infoLog:        infoLog,
		sessionManager: sessionManager,
		snippets:       snippetModel,
		templateCache:  templateCache,
		users:          userModel,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	// Initialize server with handler to access the routes available on the app
	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		TLSConfig:    tlsConfig,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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
