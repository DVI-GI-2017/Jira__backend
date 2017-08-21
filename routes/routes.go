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

func InitRouter(r *router, routes []Route) {
	for _, route := range routes {
		err := r.Route(route)
		if err != nil {
			log.Panicf("can not init route %v: %v", route, err)
		}
	}
}
