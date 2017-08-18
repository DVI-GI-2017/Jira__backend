package routes

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func InitRouter(r *router) {
	const signup= "/signup"
	err := r.Post(signup, handlers.RegisterUser)
	if err != nil {
		log.Panicf("can not init route '%s': %v", signup, err)
	}
	const signin = "/signin"
	err = r.Post(signin, handlers.Login)
	if err != nil {
		log.Panicf("can not init route '%s': %v", signin, err)
	}
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
