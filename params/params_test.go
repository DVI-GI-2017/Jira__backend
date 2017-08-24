package params

import (
	"bytes"
	"net/http"
	"net/http/httptest"
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

func TestPostBody(t *testing.T) {
	bodyBuf := bytes.NewBufferString(`{"some": ["json-like", "body"]}`)

	const path = "/"
	req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(bodyBuf.Bytes()))

	p, err := NewParams(req, regexp.MustCompile(path), path)
	if err != nil {
		t.Fail()
	}

	if !bytes.Equal(p.Body, bodyBuf.Bytes()) {
		t.Fail()
	}
}

func TestNilBody(t *testing.T) {
	const path = "/"
	req := httptest.NewRequest(http.MethodPost, path, nil)

	p, err := NewParams(req, regexp.MustCompile(path), path)
	if err != nil {
		t.Fail()
	}

	if !bytes.Equal(p.Body, nil) {
		t.Fail()
	}
}
