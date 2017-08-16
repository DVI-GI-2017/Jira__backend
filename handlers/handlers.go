package handlers

import "net/http"

func GetOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			h(w, r)

			return
		}

		http.Error(w, "Method not allowed! Get only", http.StatusMethodNotAllowed)
	}
}

func PostOnly(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			h(w, r)

			return
		}

		http.Error(w, "Method not allowed! Post only", http.StatusMethodNotAllowed)
	}
}
