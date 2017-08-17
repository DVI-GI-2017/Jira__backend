package routes

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
)

func NewRouter(rootPath string) (*router, error) {
	r := &router{}

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

type Pattern *regexp.Regexp

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
type GetHandlerFunc func(http.ResponseWriter, GetParams, PathParams)
type PostHandlerFunc func(http.ResponseWriter, PostBody, PathParams)

// Example: url "/api/v1/users/1" and pattern "/api/v1/users/:id"
// path params = {"id": "1"}
type PathParams map[string][]byte

// Get params stands for "query params"
type GetParams map[string]string

// Type for http post body
type PostBody []byte // Byte array with request body

// Add new get handler
func (r *router) Get(pattern string, handler GetHandlerFunc) error {
	compiledPattern, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	r.getHandlers[Pattern(compiledPattern)] = handler

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
		body, err := readToPostBody(req.Body)
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

// Converts url.Url.Query() from "Values" (map[string][]string)
// to "GetParams" (map[string]string)
func valuesToGetParams(values url.Values) GetParams {
	var params map[string]string
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}

// Read from reader to byte buffer
func readToPostBody(r io.ReadCloser) ([]byte, error) {
	const capacity = 100
	body := make([]byte, 0, capacity)

	_, err := r.Read(body)

	return body, err
}

func (r *router) handleGet(w http.ResponseWriter, path string, getParams GetParams) {

}

func (r *router) handlePost(w http.ResponseWriter, path string, body PostBody) {
}

func relativePath(base string, absolute string) (string, error) {
	baseLen := len(base)
	absoluteLen := len(absolute)

	if absoluteLen < baseLen {
		return "", errors.New("absolute len shorter than base len")
	}

	if absolute[:baseLen] != base {
		return "", errors.New("absolute path doesn't start with base path")
	}

	return absolute[baseLen:], nil
}
