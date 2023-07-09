package router

import (
	"fmt"
	"net/http"
)

type Middleware func(http.Handler) http.Handler

type Route struct {
	Route        string
	FinalHandler http.Handler
	// Middlewares are called in reverse order
	// (index[0](index[1](finalHandler)))
	Middleware []Middleware
}

func (route *Route) ConstructRouteHandler(handler *http.ServeMux, prefix string) {
	routePath := fmt.Sprintf("%s/%s", prefix, route.Route)

	if len(route.Middleware) > 0 {
		var handleFunc http.Handler = route.FinalHandler

		// Apply middlewares in reverse order
		for i := 0; i < len(route.Middleware); i++ {
			handleFunc = route.Middleware[i](handleFunc)
		}

		handler.Handle(routePath, handleFunc)
	} else {
		handler.Handle(routePath, route.FinalHandler)
	}
}
