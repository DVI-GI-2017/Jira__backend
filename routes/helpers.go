package routes

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strings"
)

// Type for http post body
type postBody []byte // Byte array with request body

// Get params stands for "query params"
type getParams map[string]string

// Converts url.Url.Query() from "Values" (map[string][]string)
// to "getParams" (map[string]string)
func valuesToGetParams(values url.Values) getParams {
	var params map[string]string
	for key := range values {
		params[key] = values.Get(key)
	}
	return params
}

// Example: url "/api/v1/users/1" and pattern `/api/v1/users/(?P<id>\d+)`
// path params = {"id": "1"}
type PathParams map[string]string

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

// Converts patterns like "/users/:id" to "/users/(?P<id>\d+)"
func convertSimplePatternToRegexp(pattern string) string {
	parts := strings.Split(pattern, "/")
	for i, part := range parts {
		if len(part) != 0 && part[0] == ':' {
			parts[i] = fmt.Sprintf(`(?P<%s>\d+)`, part[1:])
		}
	}

	return strings.Join(parts, "/")
}

// Return path relative to "base"
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
