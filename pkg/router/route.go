package router

import "net/http"

type Middleware func(http.Handler) http.Handler

type Route struct {
	route        string
	finalHandler http.Handler
	// Middlewares are called in reverse order
	// (index[0](index[1](finalHandler)))
	middleware []Middleware
}

func (route *Route) ConstructRouteHandler(handler *http.ServeMux) {
	if len(route.middleware) > 0 {
		var handleFunc http.Handler = route.finalHandler

		// Apply middlewares in reverse order
		for i := 0; i < len(route.middleware); i++ {
			handleFunc = route.middleware[i](handleFunc)
		}

		handler.Handle(route.route, handleFunc)
	} else {
		handler.Handle(route.route, route.finalHandler)
	}
}
