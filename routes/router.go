package routes

import (
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"net/http"
)

func NewRouter() (http.Handler, error) {

	mux := http.NewServeMux()

	mux.HandleFunc("/api/v1/login/", handlers.Login)

	return handlers.Logger(mux), nil
}
