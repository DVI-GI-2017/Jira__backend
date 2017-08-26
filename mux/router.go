package mux

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Returns new router with root path == rootPath
func NewRouter(rootPath string) (*router, error) {
	r := &router{}
	r.routes = make(routes, 0, 5)

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

// Supported methods
var supportedMethods = [...]string{
	http.MethodGet, http.MethodPost, http.MethodDelete,
	http.MethodPut, http.MethodPatch, http.MethodHead,
}

// Check if method supported
func isSupported(method string) bool {
	for _, value := range supportedMethods {
		if value == method {
			return true
		}
	}
	return false
}

type router struct {
	root *url.URL

	routes routes

	wrappers []WrapperFunc
}

// Internal structures to store routes
type route struct {
	pattern     *regexp.Regexp
	handlerFunc http.HandlerFunc
	method      string
}

type routes []route

// Set router root path, other paths will be relative to it
func (r *router) SetRootPath(path string) error {
	newRoot, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path format %s: %v", path, err)
	}

	r.root = newRoot

	return nil
}

// Add wrappers to router
func (r *router) AddWrappers(wrappers ...WrapperFunc) {
	r.wrappers = append(r.wrappers, wrappers...)
}

// Add generic route to routes.
func (r *router) Route(pattern, method string, handler http.Handler) error {
	pattern = convertSimplePatternToRegexp(pattern)

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	if !isSupported(method) {
		return fmt.Errorf("method '%s' not supported", method)
	}

	r.routes = append(r.routes, route{
		pattern:     compiledPattern,
		handlerFunc: Wrap(handler, r.wrappers...).ServeHTTP,
		method:      method,
	})

	return nil
}

// Adds Get handler
func (r *router) Get(pattern string, handler http.HandlerFunc) error {
	return r.Route(pattern, http.MethodGet, handler)
}

// Adds Post handler
func (r *router) Post(pattern string, handler http.HandlerFunc) error {
	return r.Route(pattern, http.MethodPost, handler)
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

	for _, route := range r.routes {
		if route.pattern.MatchString(path) {
			if route.method != req.Method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				fmt.Fprintf(w, "Method: '%s' not allowed on path: '%s'", req.Method, req.URL.Path)
			}

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

	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(w, "not found: %s", req.URL)
}

// Pretty prints routes
func (r *router) PrintRoutes() {
	log.Println(strings.Repeat("-", 10))

	for _, route := range r.routes {
		log.Printf("'%s': '%s'", route.method, route.pattern)
	}

	log.Println(strings.Repeat("-", 10))
}
