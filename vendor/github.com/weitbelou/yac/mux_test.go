package yac

import (
	"testing"

	"bytes"
	"net/http"
	"net/http/httptest"
)

func TestRouterCreate(t *testing.T) {
	_, err := NewRouter("/api/v1")
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestRouterResolve(t *testing.T) {
	router, _ := NewRouter("/api/v1")
	router.Get("/users/{hex:id}", func(_ http.ResponseWriter, _ *http.Request) {})

	reader := bytes.NewBufferString("")
	request := httptest.NewRequest(http.MethodGet, "/api/v1/users/1234feabc135734678123452", reader)
	response := httptest.NewRecorder()

	router.ServeHTTP(response, request)

	if response.Code == http.StatusNotFound {
		t.Fatalf("can not found route")
	}
}
