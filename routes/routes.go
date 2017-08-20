package routes

import (
	"fmt"
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
		err := initRoute(r, route)
		if err != nil {
			log.Panicf("can not init route %v: %v", route, err)
		}
	}
}

func initRoute(r *router, route Route) error {
	switch route.Method {
	case http.MethodGet:
		err := r.Get(route.Pattern, route.Handler)
		if err != nil {
			return err
		}
	case http.MethodPost:
		err := r.Post(route.Pattern, route.Handler)
		if err != nil {
			return err
		}
	default:
		return fmt.Errorf("method not supported: %s", route.Method)
	}
	return nil
}
