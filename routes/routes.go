package routes

import (
	"net/http"
	"Jira__backend/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
	Route{
		"Index",
		http.MethodGet,
		"/",
		handlers.Index(http.MethodGet),
	},
}
