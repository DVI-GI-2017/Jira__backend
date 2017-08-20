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
	var credentials models.Credentials

	parameters := params.ExtractParams(req)

	if err := json.Unmarshal(parameters.Body, &credentials); err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	action, err := pool.NewAction(pool.FindUser)
	if err != nil {
		log.Printf("%v", err)

		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "Repeat request, please!")

		return
	}

	pool.Queue <- &pool.Job{
		ModelType: credentials,
		Action:    action,
	}

	result := <-pool.Results

	if value := tools.GetValueFromModel(result.ResultType, "Email"); value == "" {
		action, err = pool.NewAction(pool.InsertUser)
		if err != nil {
			log.Printf("%v", err)

			w.WriteHeader(http.StatusBadGateway)
			fmt.Fprintln(w, "Repeat request, please!")

			return
		}

		pool.Queue <- &pool.Job{
			ModelType: credentials,
			Action:    action,
		}

		result = <-pool.Results

		tools.JsonResponse(result.ResultType, w)
	} else {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "User with this email already exists!")
	}
}

func Login(w http.ResponseWriter, req *http.Request) {
	var credentials models.Credentials

	parameters := params.ExtractParams(req)

	if err := json.Unmarshal(parameters.Body, &credentials); err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	action, err := pool.NewAction(pool.FindUser)
	if err != nil {
		log.Printf("%v", err)

		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "Repeat request, please!")

		return
	}

	pool.Queue <- &pool.Job{
		ModelType: credentials,
		Action:    action,
	}

	result := <-pool.Results

	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found!")
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
