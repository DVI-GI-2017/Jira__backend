package router

import (
	"net/http"
	"handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	NoAuth  bool // skip auth
	Handler handler
	Writer  func(handler) handlerWriter
}

type Routes []Route

func NewRouter(routes Routes) (r *pat.PatternServeMux){
	r = pat.New()

	for _, route := range routes {
		h:= route.Handler
		if !route.NoAuth{
			h = Auth()
		}

		h = HttpStatus(h)
		h = handlers.Logger(h, route.Name)
		r.Add(route.Method, route.Pattern, route.Writer(h))
	}
	return
}
