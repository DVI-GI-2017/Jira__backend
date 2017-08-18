package routes

import (
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func InitRouter(r *router) {
	r.Post("/signup", handlers.RegisterUser)
	r.Post("/signin", handlers.Login)
}
