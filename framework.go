// sword is a lightweight wrapper for Go's HTTP server made to improve the developer experience of building web services in Golang
//
// I designed this wrapper for personal use so it may be missing a certain level of polish or refinement that you would see in other projects
package sword

import (
	"net/http"
)

// function signature for HTTP handlers
type HandlerFunc = func(http.ResponseWriter, *http.Request) error

// function signature for error handling functions
type ErrorHandlerFunc = func(http.ResponseWriter, error)

// Creates a HandlerError based on a given code and message. This function is supposed to replace http.Error()
func Error(code int, message string) error {
	return HandlerError{
		code:    code,
		message: message,
	}
}
