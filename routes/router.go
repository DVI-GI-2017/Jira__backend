package routes

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"regexp"
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

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	relPath, err := relativePath(r.root.Path, req.URL.Path)
	if err != nil {
		http.NotFound(w, req)
	}

	switch method {

	case http.MethodGet:
		r.handleGet(w, req, relPath)
	case http.MethodPost:
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Panicf("invalid body: %v", err)
		}

		r.handlePost(w, req, relPath)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed: %s", method)
	}
}

func (r *router) handleGet(w http.ResponseWriter, req *http.Request, path string) {
	for pattern, route := range r.routes[http.MethodGet] {
		if pattern.MatchString(path) {
			parameters, err := params.NewParams(req, pattern, path)
			if err != nil {
				fmt.Printf("error while parsing params: %v", err)
				return
			}

			req = req.WithContext(context.WithValue(req.Context(), "params", parameters))

			route.Handler(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (r *router) handlePost(w http.ResponseWriter, req *http.Request, path string) {
	for pattern, route := range r.routes[http.MethodPost] {
		if pattern.MatchString(path) {
			parameters, err := params.NewParams(req, pattern, path)
			if err != nil {
				fmt.Printf("error while parsing params: %v", err)
				return
			}

			req = req.WithContext(context.WithValue(req.Context(), "params", parameters))

			route.Handler(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (r *router) Route(route Route) error {
	pattern := route.Pattern
	if strings.Contains(pattern, ":") {
		pattern = convertSimplePatternToRegexp(pattern)
	}

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
