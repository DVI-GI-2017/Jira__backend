package routes

import (
	"net/http"
	"Jira__backend/handlers"
)

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var RoutesList = Routes{
	Route{
		"Test",
		"/api/v1/",
		handlers.Test,
	},
}
