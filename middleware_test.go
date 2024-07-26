package sword_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grqphical/sword"

	"github.com/stretchr/testify/assert"
)

func middleware(next sword.HandlerFunc) sword.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("from middleware "))
		return next(w, r)
	}
}

func TestMiddleware(t *testing.T) {
	r := sword.NewRouter(nil)

	r.Use(middleware)

	r.RouteFunc("GET /", func(w http.ResponseWriter, r *http.Request) error {
		w.Write([]byte("from router"))
		return nil
	})

	req, err := http.NewRequest("GET", "/", nil)
	assert.NoError(t, err)

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	body := resp.Body.String()
	assert.Equal(t, 200, resp.Code)
	assert.Equal(t, "from middleware from router", body)
}
