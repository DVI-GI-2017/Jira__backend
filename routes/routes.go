package routes

import (
	"log"
	"net/http"
)

type Route struct {
	Name    string
	Pattern string
	Method  string
	Handler http.HandlerFunc
}

func (r *router) SetupRoutes() {
	InitRouter(r, LoginRoutes)
	InitRouter(r, UsersRoutes)
	InitRouter(r, ProjectRoutes)
	InitRouter(r, TasksRoutes)
	InitRouter(r, LabelsRoutes)
}

func InitRouter(r *router, routes []Route) {
	for _, route := range routes {
		err := r.Add(route)
		if err != nil {
			log.Panicf("can not init route %v: %v", route, err)
		}
	}
}
