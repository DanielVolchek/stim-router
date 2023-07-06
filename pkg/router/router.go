package router

import (
	"net/http"
)

func GetRouter(routes []*Route) *http.ServeMux {
	handler := http.NewServeMux()

	for _, route := range routes {
		route.ConstructRouteHandler(handler)
	}

	return handler
}
