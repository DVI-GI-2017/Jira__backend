package routes

import (
	"log"

	"github.com/weitbelou/yacrouter"
)

// Setup routes defined in this package
func SetupRoutes(m *mux.Router) {
	addRoutesToMux(m, defaultRoutes)
}

// Slice of default routes will be resolved automatically
var defaultRoutes mux.Routes

// Adds slice of routes to mux
func addRoutesToMux(m *mux.Router, routes mux.Routes) {
	for _, route := range routes {
		err := addRouteToMux(m, route)
		if err != nil {
			log.Panicf("can not add route %v: %v", route, err)
		}
	}
}

// Adds one route to mux
func addRouteToMux(m *mux.Router, r mux.Route) error {
	return m.Route(r.Pattern, r.Method, r.Handler)
}
