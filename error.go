package sword

import (
	"encoding/json"
	"errors"
	"net/http"
)

// represents an error returned from an HTTP handler
type HandlerError struct {
	message string
	Code    int
}

// function signature for error handling functions
type ErrorHandlerFunc = func(http.ResponseWriter, error)

// Creates a HandlerError based on a given code and message. This function is supposed to replace http.Error()
func Error(code int, message string) error {
	return HandlerError{
		Code:    code,
		message: message,
	}
}

// Ensures HandlerError derives the Error interface
func (h HandlerError) Error() string {
	return h.message
}

// sword's default error handler. It returns a JSON object with a single field 'error' which contains the error message
func defaultErrorHandler(w http.ResponseWriter, err error) {
	var handlerErr HandlerError
	if errors.As(err, &handlerErr) {
		w.WriteHeader(handlerErr.Code)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
