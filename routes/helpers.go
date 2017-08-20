package routes

import (
	"errors"
	"fmt"
	"strings"
)

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


