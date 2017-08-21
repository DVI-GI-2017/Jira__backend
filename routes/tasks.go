package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
)

var TasksRoutes = []Route{
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
