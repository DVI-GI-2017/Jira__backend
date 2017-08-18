package routes

import (
	"net/http"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/auth"
)

type RouteList []Route

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routeList = RouteList{
	Route{
		Name:        "Register new user",
		Pattern:     "/signup",
		HandlerFunc: handlers.RegisterUser,
	},
	Route{
		Name:        "Login user",
		Pattern:     "/signin",
		HandlerFunc: handlers.Login,
	},
	Route{
		Name: "All users",
		Pattern: "/users",
		HandlerFunc: auth.ValidateToken(handlers.AllUsers),
	},
	Route{
		Name: "All projects",
		Pattern: "/project",
		HandlerFunc: auth.ValidateToken(handlers.Projects),
	},
	Route{
		Name: "Project by id",
		Pattern: "/project/:id",
		HandlerFunc: auth.ValidateToken(handlers.Project),
	},
}
