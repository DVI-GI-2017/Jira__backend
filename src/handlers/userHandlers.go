package handlers

import (
	"net/http"
	"fmt"
)

func Index(method string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if method == r.Method {
			fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
		} else {
			fmt.Fprintf(w, "Bye, %s!", r.URL.Path[1:])
		}
	}
}
