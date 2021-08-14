package middleware_test

import (
	"bufio"
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/quintsys/ga-exporter/middleware"
)

func TestLogger(t *testing.T) {
	buf := &bytes.Buffer{}

	// Redirect STDOUT to the buffer
	oldStdout := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}
	os.Stdout = w
	defer func() {
		os.Stdout = oldStdout
	}()

	go func() {
		scanner := bufio.NewScanner(r)
		for scanner.Scan() {
			buf.WriteString(scanner.Text())
		}
	}()

	// Create a test HTTP server
	ts := httptest.NewServer(middleware.Logger(mockHandler()))
	defer ts.Close()

	// Make a request to get the mock handler to run and write to the buffer
	http.Get(fmt.Sprintf("%s/foo", ts.URL))
	w.Close()

	// Test buffer contains the method and url
	if !bytes.Contains(buf.Bytes(), []byte("GET /foo")) {
		t.Errorf("Expected to find 'GET /foo' in log, but didn't")
	}
}
