package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/mux"
)

func init() {
	defaultRoutes = append(defaultRoutes, projectRoutes...)
}

var projectRoutes = mux.Routes{
	{
		Pattern: "/projects",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.CreateProject),
	},
	{
		Pattern: "/projects",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllProjects),
	},
	{
		Pattern: "/projects/{hex:id}",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetProjectById),
	},
	{
		Pattern: "/projects/{hex:id}/users",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetAllUsersFromProject),
	},
	{
		Pattern: "/projects/{hex:id}/tasks",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetAllTasksFromProject),
	},
	{
		Pattern: "/projects/{hex:id}/users",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.AddUserToProject),
	},
	{
		Pattern: "/projects/{hex:id}/users/delete",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.DeleteUserFromProject),
	},
}
