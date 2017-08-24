package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func init() {
	defaultRoutes = append(defaultRoutes, tasksRoutes...)
}

var tasksRoutes = []Route{
	{
		Name:    "Create task.",
		Pattern: "/tasks",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.CreateTask),
	},
	{
		Name:    "All tasks route",
		Pattern: "/tasks",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllTasks),
	},
	{
		Name:    "Get task by id",
		Pattern: "/tasks/:id",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetTaskById),
	},
}
