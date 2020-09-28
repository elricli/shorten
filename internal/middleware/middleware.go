package middleware

import (
	"net/http"
)

// A Middleware is a func that wraps an http.Handler.
type Middleware func(http.Handler) http.Handler

// Use middleware for http handler.
func (m Middleware) Use(h http.Handler) http.Handler {
	return m(h)
}

// Chain creates a new Middleware that applies a sequence of Middlewares, so
// that they execute in the given order when handling an http request.
//
// In other words, Chain(m1, m2)(handler) = m1(m2(handler))
func Chain(middlewares ...Middleware) Middleware {
	return func(h http.Handler) http.Handler {
		for i := range middlewares {
			h = middlewares[len(middlewares)-1-i](h)
		}
		return h
	}
}
