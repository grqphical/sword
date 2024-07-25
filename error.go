package sword

import (
	"encoding/json"
	"errors"
	"net/http"
)

type HandlerError struct {
	message string
	code    int
}

func (h HandlerError) Error() string {
	return h.message
}
func defaultErrorHandler(w http.ResponseWriter, err error) {
	var handlerErr HandlerError
	if errors.As(err, &handlerErr) {
		w.WriteHeader(handlerErr.code)
	}

	json.NewEncoder(w).Encode(map[string]string{
		"error": err.Error(),
	})
}
