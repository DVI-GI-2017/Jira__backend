package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/DVI-GI-2017/Jira__backend/dataBase"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(dataBase.UsersListFromFakeDB)
}
