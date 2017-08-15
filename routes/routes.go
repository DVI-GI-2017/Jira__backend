package routes

import (
	"net/http"
	"Jira__backend/handlers"
	"Jira__backend/tools"
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
		tools.GetOnly(handlers.Test),
	},
}
