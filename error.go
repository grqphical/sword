package sword

import (
	"encoding/json"
	"errors"
	"net/http"
)

// represents an error returned from an HTTP handler
type HandlerError struct {
	message string
	code    int
}

// function signature for error handling functions
type ErrorHandlerFunc = func(http.ResponseWriter, error)

// Creates a HandlerError based on a given code and message. This function is supposed to replace http.Error()
func Error(code int, message string) error {
	return HandlerError{
		code:    code,
		message: message,
	}
}

func (h HandlerError) Error() string {
	return h.message
}

// sword's default error handler. It returns a JSON object with a single field 'error'
func defaultErrorHandler(w http.ResponseWriter, err error) {
	var handlerErr HandlerError
	if errors.As(err, &handlerErr) {
		w.WriteHeader(handlerErr.code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
