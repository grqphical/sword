package sword

import "net/http"

// configuration for the Sword router
type RouterConfig struct {
	address      string
	errorHandler ErrorHandlerFunc
}

// Router is the main wrapper for sword and handles routing, middleware, errors etc.
type Router struct {
	address      string
	mux          *http.ServeMux
	errorHandler ErrorHandlerFunc
}

func NewRouter(config *RouterConfig) *Router {
	r := &Router{}

	if config != nil {
		// a configuration was passed in
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
		// no config was provided
		r.address = ":5000"
		r.errorHandler = defaultErrorHandler
	}

	mux := http.NewServeMux()

	r.mux = mux
	return r
}

// Routes a handler function to a specific route. Sword supports the Golang net/http routing style including specifying methods and wildcards
//
// Example:
// router.RouteFunc("GET /", someHandler)
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
