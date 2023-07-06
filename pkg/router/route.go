package router

import "net/http"

type Middleware func(http.Handler) http.Handler

type Route struct {
	Route        string
	FinalHandler http.Handler
	// Middlewares are called in reverse order
	// (index[0](index[1](finalHandler)))
	Middleware []Middleware
}

func (route *Route) ConstructRouteHandler(handler *http.ServeMux) {
	if len(route.Middleware) > 0 {
		var handleFunc http.Handler = route.FinalHandler

		// Apply middlewares in reverse order
		for i := 0; i < len(route.Middleware); i++ {
			handleFunc = route.Middleware[i](handleFunc)
		}

		handler.Handle(route.Route, handleFunc)
	} else {
		handler.Handle(route.Route, route.FinalHandler)
	}
}
