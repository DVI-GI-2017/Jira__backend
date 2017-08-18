package routes

import (
	"errors"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

// Get params stands for "query params"
type getParams map[string]string

// Type for http post body
type postBody []byte // Byte array with request body

// Example: url "/api/v1/users/1" and pattern `/api/v1/users/(?P<id>\d+)`
// path params = {"id": "1"}
type PathParams map[string]string

// Converts url.Url.Query() from "Values" (map[string][]string)
// to "getParams" (map[string]string)
func valuesToGetParams(values url.Values) getParams {
	var params map[string]string
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}

// Extract path params from path
func extractPathParams(pattern *regexp.Regexp, path string) PathParams {
	match := pattern.FindStringSubmatch(path)
	result := make(PathParams)

	for i, name := range pattern.SubexpNames() {
		if i != 0 {
			result[name] = match[i]
		}
	}

	return result
}

// Read from reader to byte buffer
func readToPostBody(r io.ReadCloser) ([]byte, error) {
	const capacity = 100
	body := make([]byte, 0, capacity)

	_, err := r.Read(body)

	return body, err
}

func relativePath(base string, absolute string) (string, error) {
	baseLen := len(base)
	absoluteLen := len(absolute)

	if absoluteLen < baseLen {
		return "", errors.New("absolute len shorter than base len")
	}

	if absolute[:baseLen] != base {
		return "", errors.New("absolute path doesn't start with base path")
	}

	return absolute[baseLen:], nil
}

//Helper types for different http method handlers
type getHandlerFunc func(http.ResponseWriter, getParams, PathParams)
type postHandlerFunc func(http.ResponseWriter, postBody, PathParams)
