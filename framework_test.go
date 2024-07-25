package webframework_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	webframework "web-framework"
	responses "web-framework/responses"

	"github.com/stretchr/testify/assert"
)

func TestRouting(t *testing.T) {
	r := webframework.NewRouter(nil)

	r.RouteFunc("GET /", func(r *http.Request) (webframework.Response, error) {
		return responses.JSON(map[string]string{
			"message": "Hello, World",
		}, http.StatusOK), nil
	})

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	var data map[string]string
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&data))
	assert.Equal(t, "Hello, World", data["message"])
}
