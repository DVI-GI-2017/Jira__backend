package routes

import (
	"reflect"
	"regexp"
	"testing"

	"net/http"

	"fmt"

	"net/http/httptest"

	"bytes"

	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/gorilla/mux"
)

func TestRouter_SetRootPath(t *testing.T) {
	router := router{}

	newRoot := "/api/v1"

	err := router.SetRootPath(newRoot)

	if err != nil {
		t.Errorf("%v", err)
	}

	if router.root == nil || router.root.Path != newRoot {
		t.Errorf("%s", router.root)
	}
}

func TestRelativePath(t *testing.T) {
	const basePath = "/api/v1"
	const absolutePath = "/api/v1/users"

	relPath, err := relativePath(basePath, absolutePath)
	if err != nil {
		t.Fatal(err)
	}

	if relPath != "/users" {
		t.Fail()
	}
}

func TestExtractPathParams(t *testing.T) {
	pattern := regexp.MustCompile(`/users/(?P<id>\d+)`)
	pathParams := params.ExtractPathParams(pattern, "/users/1")

	expectedPathParams := params.PathParams{"id": "1"}

	if !reflect.DeepEqual(pathParams, expectedPathParams) {
		t.Fail()
	}
}

func TestSimplifiedPattern(t *testing.T) {
	pattern := regexp.MustCompile(convertSimplePatternToRegexp("/users/:id"))
	pathParams := params.ExtractPathParams(pattern, "/users/1")

	expectedPathParams := params.PathParams{"id": "1"}

	if !reflect.DeepEqual(pathParams, expectedPathParams) {
		t.Fail()
	}
}

func BenchmarkGorilla(b *testing.B) {
	// Create router
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api/v1").Subrouter()
	apiRouter.Path("/users/{id:[0-9]+}").Methods(http.MethodGet).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		getParams := r.URL.Query()
		vars := mux.Vars(r)
		fmt.Fprintf(w, "get params: %v, path params: %v", getParams, vars)
	})

	for i := 0; i < b.N; i++ {
		processRequest(router, b)
	}
}

func BenchmarkCustom(b *testing.B) {
	// Create router
	router, err := NewRouter("/api/v1")
	if err != nil {
		b.Errorf("can not create router: %v", err)
	}
	router.Get("/users/:id",
		func(w http.ResponseWriter, req *http.Request) {
			parameters := params.ExtractParams(req)
			fmt.Fprintf(w, "get params: %v, path params: %v", parameters.Query, parameters.PathParams)
		})

	for i := 0; i < b.N; i++ {
		processRequest(router, b)
	}
}

func processRequest(router http.Handler, b *testing.B) {
	reader := bytes.NewBufferString("")
	request := httptest.NewRequest(http.MethodGet, "/api/v1/users/1", reader)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	s := fmt.Sprintf("%s", response.Body)
	expected := "get params: map[], path params: map[id:1]"
	if s != expected {
		b.Errorf("invalid response: %s; expected: %s", s, expected)
	}
	reader = bytes.NewBufferString("")
}
