package routes

import (
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"log"
)

func loginRoutes(r *router) {
	const signup = "/signup"
	err := r.Post(signup, handlers.RegisterUser)
	if err != nil {
		log.Panicf("can not init route '%s': %v", signup, err)
	}

	const signin = "/signin"
	err = r.Post(signin, handlers.Login)
	if err != nil {
		log.Panicf("can not init route '%s': %v", signin, err)
	}
}