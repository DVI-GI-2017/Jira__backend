package main

import (
	"net/http"
	"router"
)

var routes = router.Routes{
	router.Route{
		Name:    "Login",
		Method:  "POST",
		Pattern: "/login",
		NoAuth:  true,
		Handler: LoginRequest,
		Writer:  router.JsonWriter,
	},
	router.Route{
		Name:    "Account",
		Method:  "GET",
		Pattern: "/account",
		Handler: AccountRequest,
		Writer:  router.JsonWriter,
	},
}
