package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"richwynmorris.co.uk/internal/assert"
)

func TestPing(t *testing.T) {
	app := &application{
		errorLog: log.New(io.Discard, "", 0),
		infoLog:  log.New(io.Discard, "", 0),
	}

	// Initialize a new test server to run the end to end test.
	ts := httptest.NewTLSServer(app.routes())
	defer ts.Close()

	// Generate a test client and make a request to the test server and the pin endpoint.
	result, err := ts.Client().Get(ts.URL + "/ping")
	if err != nil {
		t.Fatal(err)
	}

	// Check that result of the test has a status of 200.
	assert.Equal(t, result.StatusCode, http.StatusOK)

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}
	bytes.TrimSpace(body)

	// Check that the response contains an OK message in its body.
	assert.Equal(t, string(body), "OK")
}
