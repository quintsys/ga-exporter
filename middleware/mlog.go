package middleware

import (
	"log"
	"net/http"
)

// LogWrap adds a logger around handlers
func LogWrap(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		f(w, r)
	})
}
