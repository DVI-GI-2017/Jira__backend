package tools

import (
	"net/http"
)

func GetOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)

			return
		}

		http.Error(w, "Method not allowed! Get only", http.StatusMethodNotAllowed)
	}
}
//
//func BasicAuth(pass func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
//
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
//
//		if len(auth) != 2 || auth[0] != "Basic" {
//			http.Error(w, "authorization failed", http.StatusUnauthorized)
//			return
//		}
//
//		payload, _ := base64.StdEncoding.DecodeString(auth[1])
//		pair := strings.SplitN(string(payload), ":", 2)
//
//		if len(pair) != 2 || !validate(pair[0], pair[1]) {
//			http.Error(w, "authorization failed", http.StatusUnauthorized)
//			return
//		}
//
//		pass(w, r)
//	}
//}
//
//func validate(username, password string) bool {
//	if username == "test" && password == "test" {
//		return true
//	}
//	return false
//}
