package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

var TestRoutes = []Route{
	{
		Name:    "Test route",
		Pattern: "/test",
		Method:  http.MethodGet,
		Handler: handlers.Test,
	},
}
