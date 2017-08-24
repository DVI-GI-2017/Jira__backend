package params

import (
	"reflect"
	"regexp"
	"testing"
)

func TestExtractPathParams(t *testing.T) {
	pattern := regexp.MustCompile(`/users/(?P<id>\d+)`)
	pathParams := ExtractPathParams(pattern, "/users/12")

	expectedPathParams := PathParams{"id": "12"}

	if !reflect.DeepEqual(pathParams, expectedPathParams) {
		t.Fail()
	}
}
