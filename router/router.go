package router

import (
	"net/http"
	"Jira__backend/routes"
	"Jira__backend/handlers"
)

func NewRouter() {
	for _, route := range routes.RoutesList {
		var handler http.HandlerFunc

		handler = route.HandlerFunc
		handler = handlers.Logger(handler, route.Name)

		http.HandleFunc(route.Pattern, handler)
	}
}
