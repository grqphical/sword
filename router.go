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
	middleware   []MiddlewareFunc
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
	r.middleware = make([]MiddlewareFunc, 0)
	return r
}

// Routes a handler function to a specific route. Sword supports the Golang net/http routing style including specifying methods and wildcards
//
// Example:
// router.RouteFunc("GET /", someHandler)
func (router *Router) RouteFunc(pattern string, h HandlerFunc) {
	router.mux.HandleFunc(pattern, func(w http.ResponseWriter, r *http.Request) {
		for _, middleware := range router.middleware {
			err := middleware(h)(w, r)
			if err != nil {
				router.errorHandler(w, err)
				return
			}
		}

		if len(router.middleware) == 0 {
			err := h(w, r)
			if err != nil {
				router.errorHandler(w, err)
				return
			}
		}

	})
}

func (router *Router) Use(middleware MiddlewareFunc) {
	router.middleware = append(router.middleware, middleware)
}

// Starts the router on the defined address
func (router *Router) ListenAndServe() error {
	return http.ListenAndServe(router.address, router.mux)
}

// ensures Router satisifies the http.Handler interface
func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router.mux.ServeHTTP(w, r)
}
