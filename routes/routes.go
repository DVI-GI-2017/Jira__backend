package routes

import (
	"net/http"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
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
