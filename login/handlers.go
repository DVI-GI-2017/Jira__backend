package login

import (
	"log"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var Login = handlers.PostOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var user Credentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, "Error in request")
			log.Printf("%v", err)

			return
		}

		if err := CheckUser(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, err)
			log.Printf("%v", err)

			return
		}

		token, err := NewToken()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("%v", err)

			return
		}

		response := token
		tools.JsonResponse(response, w)
	})
