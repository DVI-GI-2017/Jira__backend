package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
)

var LabelsRoutes = []Route{
	{
		Name:    "Create label.",
		Pattern: "/labels",
		Method:  http.MethodPost,
		Handler: auth.ValidateToken(handlers.CreateLabel),
	},
	{
		Name:    "All labels route",
		Pattern: "/labels",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.AllLabels),
	},
	{
		Name:    "Get label by id",
		Pattern: "/labels/:id",
		Method:  http.MethodGet,
		Handler: auth.ValidateToken(handlers.GetLabelById),
	},
}
