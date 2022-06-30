package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/http/httptest"
	"testing"
)

func newTestApplication(t *testing.T) *application {
	return &application{
		infoLog:  log.New(io.Discard, "", 0),
		errorLog: log.New(io.Discard, "", 0),
	}
}

type testServer struct {
	*httptest.Server
}

func newTestServer(t *testing.T, h http.Handler) *testServer {
	ts := httptest.NewTLSServer(h)

	// Initialize a new cookie jar.
	jar, err := cookiejar.New(nil)
	if err != nil {
		t.Fatal(err)
	}

	// Add the newly initialized cookie jar to the client to store response cookies.
	ts.Client().Jar = jar

	// Disable the redirect on a 3XX response and instead return the immediate received response.
	ts.Client().CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return &testServer{ts}
}

func (ts *testServer) get(t *testing.T, urlPath string) (int, http.Header, string) {
	result, err := ts.Client().Get(ts.URL + urlPath)
	if err != nil {
		t.Fatal(err)
	}

	defer result.Body.Close()
	body, err := io.ReadAll(result.Body)
	if err != nil {
		t.Fatal(err)
	}

	bytes.TrimSpace(body)

	return result.StatusCode, result.Header, string(body)
}
