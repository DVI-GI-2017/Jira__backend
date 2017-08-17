package routes

import (
	"fmt"
	"net/http"
	"net/url"
)

type router struct {
	root *url.URL

	getRoutes  map[string]Route
	postRoutes map[string]Route
}

func NewRouter(rootPath string) (*router, error) {
	r := &router{}

	err := r.SetRootPath(rootPath)
	if err != nil {
		return r, err
	}

	return r, nil
}

func (r *router) SetRootPath(path string) error {
	newRoot, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path format %s: %v", path, err)
	}

	r.root = newRoot

	return nil
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	switch method {

	case http.MethodGet:
		if route, ok := r.getRoutes[path]; ok {
			route.HandlerFunc(w, req)
		}
		http.NotFound(w, req)
	case http.MethodPost:
		if route, ok := r.postRoutes[path]; ok {
			route.HandlerFunc(w, req)
		}
		http.NotFound(w, req)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed: %s", method)
	}
}
