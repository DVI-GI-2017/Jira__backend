package mux

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

var simpleBench = &benchmarkData{
	basePath: "/api/v1",

	patternGorilla: "/users/{id:[[:xdigit:]]{24}}",
	patternCustom:  "/users/:id",

	matchedPath: "/api/v1/users/599ce026ff64e74a60086508",

	pathParams:  map[string]string{"id": "599ce026ff64e74a60086508"},
	queryParams: map[string]string{},

	responseTemplate: "query params: %s, path params: %s",
}

func BenchmarkSimpleGorilla(b *testing.B) {
	benchmark(b, simpleBench, initGorillaRouter(simpleBench))
}

func BenchmarkSimpleCustom(b *testing.B) {
	benchmark(b, simpleBench, initCustomRouter(simpleBench))
}

// Executes simple benchmark atop of given data
func benchmark(b *testing.B, data *benchmarkData, router http.Handler) {
	helper := newGetHelper(data.matchedPath)
	helper.clear()

	for i := 0; i < b.N; i++ {
		router.ServeHTTP(helper.w, helper.r)
		if !data.Ok(helper.w.Body.String()) {
			b.Fail()
		}
		helper.clear()
	}
}

// Helper structure to store bench data
type benchmarkData struct {
	basePath string

	patternGorilla string
	patternCustom  string

	matchedPath string

	pathParams  map[string]string
	queryParams map[string]string

	responseTemplate string
}

func (b benchmarkData) Ok(input string) bool {
	return fmt.Sprintf(b.responseTemplate, b.queryParams, b.pathParams) == input
}

// Helpers to work with gorilla router
func initGorillaRouter(data *benchmarkData) *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix(data.basePath).Subrouter()

	apiRouter.Path(data.patternGorilla).Methods(http.MethodGet).HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			getParams := req.URL.Query()
			vars := mux.Vars(req)

			fmt.Fprintf(w, data.responseTemplate, getParams, vars)
		})

	return router
}

// Helper to work with custom router
func initCustomRouter(data *benchmarkData) *router {
	router, err := NewRouter(data.basePath)
	if err != nil {
		log.Panicf("can not create router: %v", err)
	}

	err = router.Get(data.patternCustom, func(w http.ResponseWriter, req *http.Request) {
		params := Params(req)
		fmt.Fprintf(w, data.responseTemplate, params.Query, params.PathParams)
	})
	if err != nil {
		log.Panicf("%v", err)
	}

	return router
}

// Helper to mock get requests
type getHelper struct {
	w *httptest.ResponseRecorder
	r *http.Request
}

func newGetHelper(path string) *getHelper {
	r, _ := http.NewRequest(http.MethodGet, path, nil)

	return &getHelper{
		w: httptest.NewRecorder(),
		r: r,
	}
}

func (helper *getHelper) clear() {
	helper.w.Body = bytes.NewBuffer(make([]byte, 0, 100))
}
