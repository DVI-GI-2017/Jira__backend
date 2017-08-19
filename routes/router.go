package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

func NewRouter(rootPath string) (*router, error) {
	r := &router{}
	r.getHandlers = make(map[*regexp.Regexp]getHandlerFunc)
	r.postHandlers = make(map[*regexp.Regexp]postHandlerFunc)

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

type router struct {
	root *url.URL

	getHandlers  map[*regexp.Regexp]getHandlerFunc
	postHandlers map[*regexp.Regexp]postHandlerFunc
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
		r.handleGet(w, relPath, valuesToGetParams(req.URL.Query()))
	case http.MethodPost:
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Panicf("invalid body: %v", err)
		}

		r.handlePost(w, relPath, body)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed: %s", method)
	}
}

func (r *router) handleGet(w http.ResponseWriter, path string, getParams getParams) {
	for pattern, handler := range r.getHandlers {
		if pattern.MatchString(path) {
			handler(w, getParams, extractPathParams(pattern, path))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func (r *router) handlePost(w http.ResponseWriter, path string, body postBody) {
	for pattern, handler := range r.postHandlers {
		if pattern.MatchString(path) {
			handler(w, body, extractPathParams(pattern, path))
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

// Add new GET handler
func (r *router) Get(pattern string,
	handler func(w http.ResponseWriter, getParams map[string]string, pathParams map[string]string)) error {

	if strings.Contains(pattern, ":") {
		pattern = convertSimplePatternToRegexp(pattern)
	}

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.getHandlers[compiledPattern] = func(writer http.ResponseWriter, params getParams, params2 PathParams) {
		handler(writer, params, params2)
	}

	return nil
}

// Add new POST handler
func (r *router) Post(pattern string,
	handler func(w http.ResponseWriter, body []byte, pathParams map[string]string)) error {

	if strings.Contains(pattern, ":") {
		pattern = convertSimplePatternToRegexp(pattern)
	}

	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.postHandlers[compiledPattern] = func(writer http.ResponseWriter, body postBody, params PathParams) {
		handler(writer, body, params)
	}

	return nil
}

func (r *router) Resource(resource string,
	create func(w http.ResponseWriter, body []byte, pathParams map[string]string),
	update func(w http.ResponseWriter, body []byte, pathParams map[string]string),
	receiveAll func(w http.ResponseWriter, getParams map[string]string, pathParams map[string]string),
	receiveOne func(w http.ResponseWriter, getParams map[string]string, pathParams map[string]string)) error {

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
