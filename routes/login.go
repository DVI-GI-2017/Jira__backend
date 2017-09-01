package routes

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/weitbelou/yacrouter"
)

func init() {
	defaultRoutes = append(defaultRoutes, loginRoutes...)
}

var loginRoutes = mux.Routes{
	{
		Pattern: "/signup",
		Method:  http.MethodPost,
		Handler: handlers.RegisterUser,
	},
	{
		Pattern: "/signin",
		Method:  http.MethodPost,
		Handler: handlers.Login,
	},
}
