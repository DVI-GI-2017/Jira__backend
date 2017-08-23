package routes

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"regexp"

	"log"

	"strings"

	"github.com/DVI-GI-2017/Jira__backend/params"
)

func NewRouter(rootPath string) (*router, error) {
	r := &router{}
	r.routes = make(map[string]map[*regexp.Regexp]Route)

	r.routes[http.MethodGet] = make(map[*regexp.Regexp]Route)
	r.routes[http.MethodPost] = make(map[*regexp.Regexp]Route)

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

type router struct {
	root *url.URL

	routes map[string]map[*regexp.Regexp]Route
}

// Set router root path, other paths will be relative to it
func (r *router) SetRootPath(path string) error {
	newRoot, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path format %s: %v", path, err)
	}

	r.root = newRoot

	return nil
}

// Listen on given port
func (r *router) ListenAndServe(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

// Implements http.Handler interface
func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	relPath, err := relativePath(r.root.Path, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	}

	r.handleRequest(w, req, relPath)
}

// Handles request: iterate over all routes before finds first matching route.
func (r *router) handleRequest(w http.ResponseWriter, req *http.Request, path string) {

	if routeMap, ok := r.routes[req.Method]; ok {
		for pattern, route := range routeMap {
			if pattern.MatchString(path) {
				parameters, err := params.NewParams(req, pattern, path)

				if err != nil {
					fmt.Printf("error while parsing params: %v", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				req = req.WithContext(context.WithValue(req.Context(), "params", parameters))
				route.Handler(w, req)

				return
			}
		}

		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Method: %s not allowed on path: %s", req.Method, req.URL.Path)

		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	fmt.Fprintf(w, "Method: %s not supported", req.Method)
}

// Add new route.
func (r *router) Add(route Route) error {
	pattern := route.Pattern
	pattern = convertSimplePatternToRegexp(pattern)

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	switch route.Method {
	case http.MethodGet:
		r.routes[http.MethodGet][compiledPattern] = route
		return nil
	case http.MethodPost:
		r.routes[http.MethodPost][compiledPattern] = route
		return nil
	}

	return fmt.Errorf("Error method '%s' not supported.", route.Method)
}

// Pretty prints routes
func (r *router) PrintRoutes() {
	log.Println(strings.Repeat("-", 10))

	for method, list := range r.routes {
		for re := range list {
			log.Printf("'%s': '%s'", method, re)
		}
	}

	log.Println(strings.Repeat("-", 10))
}
