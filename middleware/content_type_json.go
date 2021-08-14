package middleware

import "net/http"

// ContentTypeJSON is a middleware that sets the Content-Type header to
// application/json
func ContentTypeJSON(h http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		h.ServeHTTP(w, r)
	})
}
