// A lightweight web framework built around the Golang Standard Library HTTP server
//
// I made this as a personal tool to improve the design of my API's. The only difference this library makes is it allows for better management of middleware and
// returns errors from handlers rather than just calling `http.Error()`
package sword

import (
	"net/http"
)

// function signature for HTTP handlers
type HandlerFunc = func(http.ResponseWriter, *http.Request) error

// Wraps a Go http.Handler to a sword.HandlerFunc.
//
// Error handling for Go handlers is disabled as it is expected that Go handler's will handler errors themselves.
func WrapHandler(h http.Handler) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		h.ServeHTTP(w, r)
		return nil
	}
}

// Wraps a Go http.Handler to a sword.HandlerFunc.
//
// Error handling for Go handlers is disabled as it is expected that Go handler's will handler errors themselves.
func WrapHandlerFunc(h http.HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		h(w, r)
		return nil
	}
}
