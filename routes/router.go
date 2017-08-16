package routes

import (
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"net/http"
)

func NewRouter() (http.Handler, error) {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/users/", tools.GetOnly(tools.ValidateToken(handlers.Check)))
	mux.HandleFunc("/api/v1/login/", tools.GetOnly(handlers.Login))
	mux.HandleFunc("/api/v1/check/", tools.GetOnly(tools.ValidateToken(handlers.Check)))

	return handlers.Logger(mux), nil
}
