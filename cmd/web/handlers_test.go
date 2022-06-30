package main

import (
	"net/http"
	"testing"

	"richwynmorris.co.uk/internal/assert"
)

func TestPing(t *testing.T) {
	app := newTestApplication(t)

	// Initialize a new test server to run the end to end test.
	ts := newTestServer(t, app.routes())
	defer ts.Close()

	// Generate a test client and make a request to the test server and the ping endpoint.
	code, _, body := ts.get(t, "/ping")

	// Check that result of the test has a status of 200.
	assert.Equal(t, code, http.StatusOK)

	// Check that the response contains an OK message in its body.
	assert.Equal(t, body, "OK")
}
