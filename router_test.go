package sword_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grqphical/sword"

	"github.com/stretchr/testify/assert"
)

func TestRouting(t *testing.T) {
	r := sword.NewRouter(nil)

	r.RouteFunc("GET /", func(w http.ResponseWriter, r *http.Request) error {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("Hello, World!"))
		return nil
	})

	r.RouteFunc("GET /error", func(w http.ResponseWriter, r *http.Request) error {
		return sword.Error(http.StatusInternalServerError, "error")
	})

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body := resp.Body.String()
	assert.Equal(t, "Hello, World!", body)
	assert.Equal(t, 202, resp.Code)

	resp = httptest.NewRecorder()

	req, err = http.NewRequest("GET", "/error", nil)
	assert.NoError(t, err)
	r.ServeHTTP(resp, req)

	var respData map[string]string
	assert.NoError(t, json.NewDecoder(resp.Body).Decode(&respData))
	assert.Equal(t, "error", respData["error"])
	assert.Equal(t, 500, resp.Code)
}
