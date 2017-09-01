package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/weitbelou/yac"
)

func init() {
	defaultRoutes = append(defaultRoutes, labelsRoutes...)
}

var labelsRoutes = yac.Routes{
	{
		Pattern: "/tasks/{hex:task_id}/labels",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.AddLabelToTask),
	},
	{
		Pattern: "/tasks/{hex:task_id}/labels",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllLabelsOnTask),
	},
	{
		Pattern: "/tasks/{hex:task_id}/labels/delete",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.DeleteLabelFromTask),
	},
}
