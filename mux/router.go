package mux

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func NewRouter(rootPath string) (*router, error) {
	r := &router{}
	r.routes = make(map[string]map[*regexp.Regexp]http.HandlerFunc)

	r.routes[http.MethodGet] = make(map[*regexp.Regexp]http.HandlerFunc)
	r.routes[http.MethodPost] = make(map[*regexp.Regexp]http.HandlerFunc)

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

type router struct {
	root *url.URL

	routes map[string]map[*regexp.Regexp]http.HandlerFunc

	wrappers []WrapperFunc
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

// Add wrappers to router
func (r *router) AddWrappers(wrappers ...WrapperFunc) {
	r.wrappers = append(r.wrappers, wrappers...)
}

// Adds Get handler
func (r *router) Get(pattern string, handler http.HandlerFunc) error {
	pattern = convertSimplePatternToRegexp(pattern)

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.routes[http.MethodGet][compiledPattern] = Wrap(handler, r.wrappers...).ServeHTTP

	return nil
}

// Adds Post handler
func (r *router) Post(pattern string, handler http.HandlerFunc) error {
	pattern = convertSimplePatternToRegexp(pattern)

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.routes[http.MethodPost][compiledPattern] = Wrap(handler, r.wrappers...).ServeHTTP

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
		for pattern, handler := range routeMap {
			if pattern.MatchString(path) {
				params, err := newParams(req, pattern, path)

				if err != nil {
					fmt.Printf("error while parsing params: %v", err)
					w.WriteHeader(http.StatusBadRequest)
					return
				}

				handler.ServeHTTP(w, putParams(req, params))

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
func (r *router) PrintRoutes() {
	log.Println(strings.Repeat("-", 10))

	for method, list := range r.routes {
		for re := range list {
			log.Printf("'%s': '%s'", method, re)
		}
	}

	log.Println(strings.Repeat("-", 10))
}
