package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
)

var UsersRoutes = []Route{
	{
		Name:    "All users route",
		Pattern: "/users",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllUsers),
	},
}
