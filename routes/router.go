package routes

import (
	"fmt"
	"net/http"
	"net/url"
)

func NewRouter(rootPath string) (*router, error) {
	r := &router{}

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

type Pattern string

type router struct {
	root *url.URL

	getHandlers  map[Pattern]GetHandlerFunc
	postHandlers map[Pattern]PostHandlerFunc
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

//Helper types for different http method handlers
type GetHandlerFunc func(*http.ResponseWriter, GetParams, PathParams)
type PostHandlerFunc func(*http.ResponseWriter, PostBody, PathParams)

// Example: url "/api/v1/users/1" and pattern "/api/v1/users/:id"
// path params = {"id": "1"}
type PathParams map[string][]byte

// Get params stands for "query params"
type GetParams map[string]string

// Converts url.Url.Query() from "Values" (map[string][]string)
// to "GetParams" (map[string]string)
func valuesToGetParams(values url.Values) GetParams {
	var params map[string]string
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}

// Type for http post body
type PostBody []byte // Byte array with request body

// Add new get handler
func (r *router) Get(pattern string, handler GetHandlerFunc) error {
	fullPattern, err := r.root.Parse(pattern)
	if err != nil {
		return err
	}

	r.getHandlers[Pattern(fullPattern.Path)] = handler

	return nil
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	method := req.Method

	switch method {

	case http.MethodGet:
		r.handleGet(w, req)
	case http.MethodPost:
		r.handlePost(w, req)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed: %s", method)
	}
}

func (r *router) handlePost(w http.ResponseWriter, req *http.Request) {

}

func (r *router) handleGet(w http.ResponseWriter, req *http.Request) {

}
