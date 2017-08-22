package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/services/auth"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	credentials := new(models.User)

	parameters := params.ExtractParams(req)

	if err := json.Unmarshal(parameters.Body, credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	exists, err := pool.DispatchAction(pool.CheckUserExists, credentials)
	if err != nil {
		log.Panicf("%v", err)
	}

	if exists.(bool) {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "User with email: %s already exists!", credentials.Email)

		log.Printf("User with email: %s already exists!", credentials.Email)

		return
	}

	user, err := pool.DispatchAction(pool.CreateUser, credentials)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "Can not create your account. Please, try later")
		log.Printf("can not create user: %v", err)

		return
	}

	tools.JsonResponse(user, w)
}

func Login(w http.ResponseWriter, req *http.Request) {
	credentials := new(models.User)

	parameters := params.ExtractParams(req)

	if err := json.Unmarshal(parameters.Body, credentials); err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	valid, err := pool.DispatchAction(pool.CheckUserCredentials, credentials)
	if err != nil {
		log.Panicf("%v", err)
	}

	if !valid.(bool) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found!")
		return
	}

	token, err := auth.NewToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		fmt.Fprintln(w, "Error while signing the token!")
		log.Printf("%v", err)

		return
	}

	tools.JsonResponse(token, w)
}
