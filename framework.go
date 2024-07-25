package sword

import (
	"net/http"
)

type HandlerFunc = func(http.ResponseWriter, *http.Request) error
type ErrorHandlerFunc = func(http.ResponseWriter, error)

func Error(code int, message string) error {
	return HandlerError{
		code:    code,
		message: message,
	}
}
