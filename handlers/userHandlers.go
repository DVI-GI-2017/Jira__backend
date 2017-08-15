package handlers

import (
	"net/http"
	"fmt"
	"encoding/json"
	"Jira__backend/dataBase"
)

func Test(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		json.NewEncoder(w).Encode(dataBase.UsersListFromFakeDB)
	default:
		w.WriteHeader(200)
		fmt.Fprintln(w, "Method not allowed!")
	}
}
