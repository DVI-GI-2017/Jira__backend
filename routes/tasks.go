package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/weitbelou/yacrouter"
)

func init() {
	defaultRoutes = append(defaultRoutes, tasksRoutes...)
}

var tasksRoutes = mux.Routes{
	{
		Pattern: "/projects/{hex:project_id}/tasks",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.AddTaskToProject),
	},
	{
		Pattern: "/projects/{hex:project_id}/tasks",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllTasksInProject),
	},
	{
		Pattern: "/tasks/{hex:id}",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetTaskById),
	},
}
