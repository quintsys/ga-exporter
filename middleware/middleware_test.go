package middleware_test

import (
	"net/http"
	"strconv"
	"testing"

	"github.com/quintsys/ga-exporter/middleware"
)

func TestChainMiddlewares(t *testing.T) {
	// create an array of 3 middlewares
	var middlewares []middleware.Middleware
	var called []string
	for i := 0; i < 3; i++ {
		order := i + 1
		m := func(h http.HandlerFunc) http.HandlerFunc {
			return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				called = append(called, strconv.Itoa(order))
				h.ServeHTTP(w, r)
			})
		}
		middlewares = append(middlewares, m)
	}

	// mock the original http.HandlerFunc
	mockHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		called = append(called, "mock")
	})

	// call the middleware chainer
	middleware.ChainMiddlewares(mockHandler, middlewares...)(nil, nil)

	// check that middlewares were called in the correct order
	for i := 0; i < len(called)-1; i++ {
		if called[i] != strconv.Itoa(i+1) {
			t.Errorf("Expected %s, got %s", strconv.Itoa(i+1), called[i])
		}
	}

	// check that the mock handler was called last
	last := called[len(called)-1]
	if last != "mock" {
		t.Errorf("Expected %s, got %s", "mock", last)
	}
}

func mockHandler() http.HandlerFunc {
	fn := func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}
	return http.HandlerFunc(fn)
}
