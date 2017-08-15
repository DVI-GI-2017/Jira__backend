package routes

import (
	"net/http"
	"Jira__backend/tools"
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
		tools.GetOnly(handlers.Test),
	},
	Route{
		"Test",
		"/api/v1/login",
		tools.PostOnly(handlers.Login),
	},
	Route{
		"Test",
		"/api/v1/check",
		tools.GetOnly(tools.ValidateToken(handlers.Check)),
	},
}
