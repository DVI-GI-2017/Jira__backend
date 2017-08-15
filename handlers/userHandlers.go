package handlers

import (
	"fmt"
	"net/http"
	"encoding/json"
	"Jira__backend/dataBase"
	"Jira__backend/models"
	"Jira__backend/validators"
	"Jira__backend/tools"
)

func Test(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataBase.UsersListFromFakeDB)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	if err := validators.CheckUser(user); err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, err)
		return
	}

	tokenString, err := tools.CreateToken()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")

		return
	}

	response := tools.Token{Token: tokenString}
	tools.JsonResponse(response, w)
}

func Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataBase.UsersListFromFakeDB)
}
