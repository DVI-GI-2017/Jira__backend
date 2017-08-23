package mux

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Converts patterns like "/users/:id" to "/users/(?P<id>\d+)"
func convertSimplePatternToRegexp(pattern string) string {
	parts := strings.Split(pattern, "/")

	for i, part := range parts {
		if len(part) != 0 && part[0] == ':' {
			parts[i] = fmt.Sprintf(`(?P<%s>[a-f\d]{24})`, part[1:])
		}
	}

	pattern = strings.Join(parts, `\/`)
	pattern = fmt.Sprintf("^%s$", pattern)

	return pattern
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

// Logs requests
func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		handler.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	})
}
