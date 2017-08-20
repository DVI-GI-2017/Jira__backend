package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"log"
	"net/http"
)

func RegisterUser(w http.ResponseWriter, body []byte, _ map[string]string) {
	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	action, err := pool.NewAction(pool.Insert)
	if err != nil {
		log.Printf("%v", err)

		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "Repeat, please!")

		return
	}

	pool.Queue <- &pool.Job{
		ModelType: user,
		Action:    action,
	}

	result := <-pool.Results

	tools.JsonResponse(result.ResultType, w)
}

func Login(w http.ResponseWriter, body []byte, _ map[string]string) {
	var user models.User

	if err := json.Unmarshal(body, &user); err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	action, err := pool.NewAction(pool.Find)
	if err != nil {
		log.Printf("%v", err)

		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "Repeat, please!")

		return
	}

	pool.Queue <- &pool.Job{
		ModelType: user,
		Action:    action,
	}

	result := <-pool.Results

	if result.Error != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Unauthorized!")
	} else {
		token, err := auth.NewToken()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, "Error while signing the token!")
			log.Printf("%v", err)

			return
		}

		tools.JsonResponse(token, w)
	}
}
