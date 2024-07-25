package webframework

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	data   any
	status int
}

func JSON(data any, status int) JSONResponse {
	return JSONResponse{
		data,
		status,
	}
}

func (j JSONResponse) WriteResponse(w http.ResponseWriter) error {
	w.WriteHeader(j.status)
	return json.NewEncoder(w).Encode(j.data)
}
