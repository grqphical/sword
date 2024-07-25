package webframework

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

type HandlerFunc = func(http.ResponseWriter, *http.Request) error
type ErrorHandlerFunc = func(http.ResponseWriter, error)

type RouterConfig struct {
	address      string
	errorHandler ErrorHandlerFunc
}

type Router struct {
	address      string
	mux          *http.ServeMux
	errorHandler ErrorHandlerFunc
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

func Error(code int, message string) error {
	return HandlerError{
		code:    code,
		message: message,
	}
}

func NewRouter(config *RouterConfig) *Router {
	r := &Router{}

	if config != nil {
		if config.address == "" {
			r.address = ":5000"
		} else {
			r.address = config.address
		}

		if config.errorHandler == nil {
			r.errorHandler = defaultErrorHandler
		} else {
			r.errorHandler = config.errorHandler
		}
	} else {
		r.address = ":5000"
		r.errorHandler = defaultErrorHandler
	}

	mux := http.NewServeMux()

	r.mux = mux
	return r
}

func (router *Router) RouteFunc(pattern string, h HandlerFunc) {
	router.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		if err != nil {
			router.errorHandler(w, err)
			return
		}
	})
}

func (router *Router) ListenAndServe() error {
	return http.ListenAndServe(router.address, router.mux)
}

func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
