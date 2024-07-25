package sword

type MiddlewareFunc = func(HandlerFunc) HandlerFunc
