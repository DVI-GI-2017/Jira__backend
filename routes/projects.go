package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

var ProjectRoutes = []Route{
	{
		Name:    "Creates project",
		Pattern: "/projects",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.CreateProject),
	},
	{
		Name:    "Get all projects",
		Pattern: "/projects",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllProjects),
	},
	{
		Name:    "Get project by id",
		Pattern: "/projects/:id",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetProjectById),
	},
}
