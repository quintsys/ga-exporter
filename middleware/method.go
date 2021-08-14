package middleware

import (
	"net/http"
)

// PostOnly is a middleware that only allows POST requests
func PostOnly(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			h.ServeHTTP(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
