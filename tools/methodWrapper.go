package tools

import "net/http"

func GetOnly(h func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			h(w, r)

			return
		}

		http.Error(w, "Method not allowed! Get only", http.StatusMethodNotAllowed)
	}
}
