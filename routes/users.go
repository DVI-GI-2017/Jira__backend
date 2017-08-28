package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/mux"
)

func init() {
	defaultRoutes = append(defaultRoutes, usersRoutes...)
}

var usersRoutes = mux.Routes{
	{
		Pattern: "/users/{hex:id}",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetUserById),
	},
	{
		Pattern: "/users",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllUsers),
	},
	{
		Pattern: "/users/{hex:id}/projects",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetAllProjectsFromUser),
	},
	{
		Pattern: "/cur-user",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.JsonNullHandler),
	},
}
