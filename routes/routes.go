package routes

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

func InitRouter(r *router) {
	//const signup = "/signup"
	//err := r.Post(signup, handlers.RegisterUser)
	//if err != nil {
	//	log.Panicf("can not init route '%s': %v", signup, err)
	//}
	//
	//const signin = "/signin"
	//err = r.Post(signin, handlers.Login)
	//if err != nil {
	//	log.Panicf("can not init route '%s': %v", signin, err)
	//}

	const test = "/test"
	err := r.Post(test, handlers.Test)
	if err != nil {
		log.Panicf("can not init route '%s': %v", test, err)
	}
}
