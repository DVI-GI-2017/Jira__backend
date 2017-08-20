package routes

import (
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"log"
)

func testRoutes(r *router) {
	const test = "/test"
	err := r.Post(test, handlers.Test)
	if err != nil {
		log.Panicf("can not init route '%s': %v", test, err)
	}
}
