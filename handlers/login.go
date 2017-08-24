package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
)

// Registers user
// Post body - user credentials in format: {"email": "...", "password": "..."}
// Returns credentials if OK
func RegisterUser(w http.ResponseWriter, req *http.Request) {
	var credentials models.User

	params := mux.Params(req)

	if err := json.Unmarshal(params.Body, &credentials); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := credentials.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.DispatchAction(pool.CheckUserExists, credentials)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("User with email: %s already exists!", credentials.Email),
			http.StatusConflict)
		return
	}

	user, err := pool.DispatchAction(pool.CreateUser, credentials)
	if err != nil {
		JsonErrorResponse(w, fmt.Errorf("can not create account: %v", err), http.StatusBadGateway)
		return
	}

	token, err := auth.NewToken()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, struct {
		models.User
		auth.Token
	}{user.(models.User), token})
}

// Authorizes user in system.
// Post body - credentials in format: {"email": "...", "password": "..."}
// Returns token for authentication.
func Login(w http.ResponseWriter, req *http.Request) {
	var credentials models.User

	params := mux.Params(req)

	if err := json.Unmarshal(params.Body, &credentials); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := credentials.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	valid, err := pool.DispatchAction(pool.CheckUserCredentials, credentials)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if !valid.(bool) {
		JsonErrorResponse(w, fmt.Errorf("can not find user with: %v", credentials), http.StatusNotFound)
		return
	}

	token, err := auth.NewToken()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	JsonResponse(w, token)
}
