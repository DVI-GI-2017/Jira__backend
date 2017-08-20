package routes

import (
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"log"
)

func usersRoutes(r *router) {
	const users = "/users"
	err := r.Get(users, handlers.AllUsers)
	if err != nil {
		log.Panicf("can not init route '%s': %v", users, err)
	}
}
