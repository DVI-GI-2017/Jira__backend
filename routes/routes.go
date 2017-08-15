package routes

import (
	"net/http"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/tools"
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
