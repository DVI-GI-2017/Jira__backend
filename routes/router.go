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
	r.getHandlers = make(map[*regexp.Regexp]http.HandlerFunc)
	r.postHandlers = make(map[*regexp.Regexp]http.HandlerFunc)

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

type router struct {
	root *url.URL

	getHandlers  map[*regexp.Regexp]http.HandlerFunc
	postHandlers map[*regexp.Regexp]http.HandlerFunc
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
	for pattern, handler := range r.getHandlers {
		if pattern.MatchString(path) {
			parameters, err := params.NewParams(req, pattern, path)
			if err != nil {
				fmt.Printf("error while parsing params: %v", err)
				return
			}

			req = req.WithContext(context.WithValue(req.Context(), "params", parameters))

			handler(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (r *router) handlePost(w http.ResponseWriter, req *http.Request, path string) {
	for pattern, handler := range r.postHandlers {
		if pattern.MatchString(path) {
			parameters, err := params.NewParams(req, pattern, path)
			if err != nil {
				fmt.Printf("error while parsing params: %v", err)
				return
			}

			req = req.WithContext(context.WithValue(req.Context(), "params", parameters))

			handler(w, req)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Add new GET handler
func (r *router) Get(pattern string, handler http.HandlerFunc) error {

	if strings.Contains(pattern, ":") {
		pattern = convertSimplePatternToRegexp(pattern)
	}

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.getHandlers[compiledPattern] = handler

	return nil
}

// Add new POST handler
func (r *router) Post(pattern string, handler http.HandlerFunc) error {

	if strings.Contains(pattern, ":") {
		pattern = convertSimplePatternToRegexp(pattern)
	}

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.postHandlers[compiledPattern] = handler

	return nil
}

func (r *router) Resource(resource string, create, update, receiveAll, receiveOne http.HandlerFunc) error {

	resourceById := fmt.Sprintf("%s/:id", resource)

	if err := r.Post(resource, create); err != nil {
		return fmt.Errorf("can not init 'create' route: %v", err)
	}

	if err := r.Get(resource, receiveAll); err != nil {
		return fmt.Errorf("can not init 'receive all' route: %v", err)
	}

	if err := r.Get(resourceById, receiveOne); err != nil {
		return fmt.Errorf("can not init 'receive one' route: %v", err)
	}

	if err := r.Post(resourceById, update); err != nil {
		return fmt.Errorf("can not init 'update' route: %v", err)
	}

	return nil
}
