package routes

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

type router struct {
	root *url.URL

	getHandlers  map[string]GetHandler
	postHandlers map[string]PostHandler
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
		if route, ok := r.getHandlers[path]; ok {
			route.HandleGet(&w, valuesToGetParams(req.URL.Query()), nil)
		}
		http.NotFound(w, req)
	case http.MethodPost:
		if route, ok := r.postHandlers[path]; ok {
			var body []byte
			_, err := req.Body.Read(body)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				log.Panicf("can not read request body: %v", err)
			}

			route.HandlePost(&w, RequestBody(body), nil)
		}
		http.NotFound(w, req)

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "method not allowed: %s", method)
	}
}

type GetParams map[string]string
type PathParams map[string][]byte

type GetHandler interface {
	HandleGet(*http.ResponseWriter, GetParams, PathParams)
}

type RequestBody []byte // Byte array with request body

type PostHandler interface {
	HandlePost(*http.ResponseWriter, RequestBody, PathParams)
}

func valuesToGetParams(values url.Values) GetParams {
	var params map[string]string
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}
