package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
)

var LabelsRoutes = []Route{
	{
		Name:    "Add label to task (:task_id)",
		Pattern: "/tasks/:task_id/labels",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.AddLabelToTask),
	},
	{
		Name:    "All labels on task",
		Pattern: "/tasks/:task_id/labels",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllLabelsOnTask),
	},
}
