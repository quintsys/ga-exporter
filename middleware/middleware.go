package middleware

import "net/http"

// Middleware is a function that can be used as a middleware
type Middleware func(next http.HandlerFunc) http.HandlerFunc

// ChainMiddlewares chains middlewares together and returns a new handler
func ChainMiddlewares(h http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for i := len(middlewares) - 1; i >= 0; i-- {
		h = middlewares[i](h)
	}
	return h
}
