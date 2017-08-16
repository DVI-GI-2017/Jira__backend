package routes

import (
	"net/http"
)

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
