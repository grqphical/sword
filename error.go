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

func (h HandlerError) Error() string {
	return h.message
}

// sword's default error handler. It returns a JSON object with a single field 'error'
func defaultErrorHandler(w http.ResponseWriter, err error) {
	var handlerErr HandlerError
	if errors.As(err, &handlerErr) {
		w.WriteHeader(handlerErr.code)
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
