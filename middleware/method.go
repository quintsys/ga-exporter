package middleware

import (
	"net/http"
)

// PostOnly allows POST requests only
func PostOnly(f func(w http.ResponseWriter, r *http.Request)) func(
	http.ResponseWriter, *http.Request) {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "POST":
			f(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

}
