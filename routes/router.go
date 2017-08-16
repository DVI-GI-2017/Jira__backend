package routes

import (
	"fmt"
	"net/http"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"github.com/DVI-GI-2017/Jira__backend/login"
)

func NewRouter() (http.Handler, error) {
	const apiRoot = "/api/v1"

	mux := http.NewServeMux()

	mux.HandleFunc(fmt.Sprintf("%s%s", apiRoot, "/login/"), login.Login)

	return handlers.Logger(mux), nil
}

type Router struct {
	mux http.ServeMux
}
