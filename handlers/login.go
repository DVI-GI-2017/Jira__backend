package handlers

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"github.com/DVI-GI-2017/Jira__backend/login"
)

var Login = PostOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var user login.Credentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, "Error in request")
			log.Printf("%v", err)

			return
		}

		if err := login.LoginUser(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, err)
			log.Printf("%v", err)

			return
		}

		token, err := login.NewToken()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("%v", err)

			return
		}

		response := token
		tools.JsonResponse(response, w)
	})
