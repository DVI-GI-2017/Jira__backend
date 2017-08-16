package routes

import (
	"net/http"
)

type RouteList []Route

type Route struct {
	Name        string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var routeList = RouteList{

}
