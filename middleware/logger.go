package middleware

import (
	"log"
	"net/http"
	"os"
)

// Logger adds a logger around handlers
func Logger(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.SetOutput(os.Stdout)
		log.Printf("%s %s", r.Method, r.URL.Path)

		// call the original handler
		h.ServeHTTP(w, r)
	})
}
