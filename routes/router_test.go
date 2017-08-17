package routes

import (
	"reflect"
	"regexp"
	"testing"
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
	pathParams := extractPathParams(pattern, "/users/1")

	expectedPathParams := PathParams{"id": "1"}

	if !reflect.DeepEqual(pathParams, expectedPathParams) {
		t.Fail()
	}
}
