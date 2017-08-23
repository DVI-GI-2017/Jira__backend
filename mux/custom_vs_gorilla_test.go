package mux

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/gorilla/mux"
)

func BenchmarkGorilla(b *testing.B) {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.Path("/users/{id:[a-f0-9]{24}}").Methods(http.MethodGet).HandlerFunc(gorillaHandler)

	for i := 0; i < b.N; i++ {
		processRequest(router, b)
	}
}

func gorillaHandler(w http.ResponseWriter, req *http.Request) {
	getParams := req.URL.Query()
	vars := mux.Vars(req)

	fmt.Fprintf(w, "get params: %v, path params: %v", getParams, vars)
}

func BenchmarkCustom(b *testing.B) {
	router, err := NewRouter("/api/v1")
	if err != nil {
		b.Errorf("can not create router: %v", err)
	}

	err = router.Get("/users/:id", customHandler)
	if err != nil {
		b.Fatalf("%v", err)
	}

	for i := 0; i < b.N; i++ {
		processRequest(router, b)
	}
}

func customHandler(w http.ResponseWriter, req *http.Request) {
	parameters := params.ExtractParams(req)
	fmt.Fprintf(w, "get params: %v, path params: %v", parameters.Query, parameters.PathParams)
}

func processRequest(router http.Handler, b *testing.B) {
	reader := bytes.NewBufferString("")
	request := httptest.NewRequest(http.MethodGet, "/api/v1/users/234feabc1357346781234524", reader)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	s := fmt.Sprintf("%s", response.Body)
	expected := "get params: map[], path params: map[id:234feabc1357346781234524]"
	if s != expected {
		b.Errorf("invalid response: %s; expected: %s", s, expected)
	}
}
