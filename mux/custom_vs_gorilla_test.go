package mux

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func BenchmarkGorilla(b *testing.B) {
	router := mux.NewRouter()

	apiRouter := router.PathPrefix("/api/v1").Subrouter()

	apiRouter.Path("/users/{id:[a-f0-9]{24}}").Methods(http.MethodGet).HandlerFunc(gorillaHandler)

	helper := newGetHelper("/api/v1/users/234feabc1357346781234524")

	for i := 0; i < b.N; i++ {
		processRequest(router, helper, b)
		helper.clear()
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

	helper := newGetHelper("/api/v1/users/234feabc1357346781234524")

	for i := 0; i < b.N; i++ {
		processRequest(router, helper, b)
		helper.clear()
	}
}

func customHandler(w http.ResponseWriter, req *http.Request) {
	params := Params(req)
	fmt.Fprintf(w, "get params: %v, path params: %v", params.Query, params.PathParams)
}

func processRequest(router http.Handler, helper *getHelper, b *testing.B) {
	router.ServeHTTP(helper.w, helper.r)

	s := fmt.Sprintf("%s", helper.w.Body)
	expected := "get params: map[], path params: map[id:234feabc1357346781234524]"
	if s != expected {
		b.Errorf("invalid response: %s; expected: %s", s, expected)
	}
}

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
