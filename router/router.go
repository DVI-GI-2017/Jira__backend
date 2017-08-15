package router

import (
	"net/http"
	"github.com/DVI-GI-2017/Jira__backend/routes"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func NewRouter() {
	for _, route := range routes.RoutesList {
		var handler http.HandlerFunc

		handler = route.HandlerFunc
		handler = handlers.Logger(handler, route.Name)

		http.HandleFunc(route.Pattern, handler)
	}
}
