// sword is a lightweight wrapper for Go's HTTP server made to improve the developer experience of building web services in Golang
//
// I designed this wrapper for personal use so it may be missing a certain level of polish or refinement that you would see in other projects
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
