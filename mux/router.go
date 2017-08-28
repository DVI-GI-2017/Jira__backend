package mux

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Helpers for simpler usage
type Pattern string
type Method string

// Structure to store routes
type Route struct {
	pattern     *regexp.Regexp
	handlerFunc http.HandlerFunc
}

// Structure that stores supported methods for each route.
type Routes map[Pattern]map[Method]Route

// Returns new router with root path == rootPath
func NewRouter(root string) (*Router, error) {
	r := &Router{}
	r.routes = make(Routes)

	newRoot, err := url.Parse(root)
	if err != nil {
		return nil, fmt.Errorf("invalid path format %s: %v", root, err)
	}

	r.root = newRoot

	return r, nil
}

type Router struct {
	root *url.URL

	routes map[Pattern]map[Method]Route

	wrappers []WrapperFunc
}

// Add wrappers to router
func (r *Router) AddWrappers(wrappers ...WrapperFunc) {
	r.wrappers = append(r.wrappers, wrappers...)
}

// Add generic route to routes.
func (r *Router) Route(pattern, method string, handler http.Handler) error {
	pattern = convertSimplePatternToRegexp(pattern)

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if _, ok := r.routes[method]; !ok {
		return fmt.Errorf("method '%s' not supported", method)
	}

	r.routes[method] = append(r.routes[method],
		Route{compiledPattern, Wrap(handler, r.wrappers...).ServeHTTP})

	return nil
}

// Adds Get handler
func (r *Router) Get(pattern string, handler http.HandlerFunc) error {
	return r.Route(pattern, http.MethodGet, handler)
}

// Adds Post handler
func (r *Router) Post(pattern string, handler http.HandlerFunc) error {
	return r.Route(pattern, http.MethodPost, handler)
}

// Listen on given port
func (r *Router) ListenAndServe(port string) error {
	return http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}

// Implements http.Handler interface
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	relPath, err := relativePath(r.root.Path, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	}

	r.handleRequest(w, req, relPath)
}

// Handles request: iterate over all routes before finds first matching route.
func (r *Router) handleRequest(w http.ResponseWriter, req *http.Request, path string) {

	if routes, ok := r.routes[req.Method]; ok {
		for _, route := range routes {
			if route.pattern.MatchString(path) {
				params, err := newParams(req, route.pattern, path)

				if err != nil {
					fmt.Printf("error while parsing params: %v", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				route.handlerFunc(w, putParams(req, params))

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

// Pretty prints routes
func (r *Router) PrintRoutes() {
	log.Println(strings.Repeat("-", 10))

	for method, list := range r.routes {
		for _, r := range list {
			log.Printf("'%s': '%s'", method, r.pattern)
		}
	}

	log.Println(strings.Repeat("-", 10))
}
