package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/services"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"log"
	"net/http"
)

var RegisterUser = PostOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var credentials auth.Credentials

		if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
			w.WriteHeader(http.StatusBadRequest)

			fmt.Fprint(w, "Invalid registration data format.")
			log.Printf("%v", err)

			return
		}

		if _, err := services.GetUserByEmailAndPassword(credentials.Email, credentials.Password); err == nil {
			w.WriteHeader(http.StatusConflict)

			fmt.Fprint(w, "User with this email already exists.")

			return
		}

		err := services.AddUser(&credentials)

		if err != nil {
			fmt.Fprint(w, "Error insert")
			w.WriteHeader(http.StatusBadGateway)

			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, credentials)
	})

var Login = PostOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var user auth.Credentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, "Error in request")
			log.Printf("%v", err)

			return
		}

		if err := auth.LoginUser(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, err)
			log.Printf("%v", err)

			return
		}

		token, err := auth.NewToken()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("%v", err)

			return
		}

		response := token
		tools.JsonResponse(response, w)
	})
