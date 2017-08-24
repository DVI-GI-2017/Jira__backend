package mux

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/DVI-GI-2017/Jira__backend/params"
)

func TestSimplifiedPattern(t *testing.T) {
	pattern := regexp.MustCompile(convertSimplePatternToRegexp("/users/:id"))
	pathParams := params.ExtractPathParams(pattern, "/users/234feabc1357346781234524")

	expectedPathParams := params.PathParams{"id": "234feabc1357346781234524"}

	if !reflect.DeepEqual(pathParams, expectedPathParams) {
		t.Fatalf("expected: %v but got: %v", expectedPathParams, pathParams)
	}
}

func TestRelativePath(t *testing.T) {
	const basePath = "/api/v1"
	const absolutePath = "/api/v1/users/1234feabc1357346781234524"

	relPath, err := relativePath(basePath, absolutePath)
	if err != nil {
		t.Fatal(err)
	}

	if relPath != "/users/1234feabc1357346781234524" {
		t.Fail()
	}
}
