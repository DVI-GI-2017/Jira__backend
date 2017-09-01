package routes

import (
	"log"

	"github.com/weitbelou/yac"
)

// Setup routes defined in this package
func SetupRoutes(m *yac.Router) {
	addRoutesToMux(m, defaultRoutes)
}

// Slice of default routes will be resolved automatically
var defaultRoutes yac.Routes

// Adds slice of routes to mux
func addRoutesToMux(m *yac.Router, routes yac.Routes) {
	for _, route := range routes {
		err := addRouteToMux(m, route)
		if err != nil {
			log.Panicf("can not add route %v: %v", route, err)
		}
	}
}

// Adds one route to mux
func addRouteToMux(m *yac.Router, r yac.Route) error {
	return m.Route(r.Pattern, r.Method, r.Handler)
}
