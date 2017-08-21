package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

func Test(w http.ResponseWriter, req *http.Request) {
	var user models.User

	parameters := params.ExtractParams(req)

	if err := json.Unmarshal(parameters.Body, &user); err != nil {
		w.WriteHeader(http.StatusForbidden)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	action, err := pool.NewAction(pool.FindUserById)
	if err != nil {
		log.Printf("%v", err)

		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprintln(w, "Repeat, please!")

		return
	}

	pool.Queue <- &pool.Job{
		Input:  user,
		Action: action,
	}

	result := <-pool.Results

	if value := tools.GetValueFromModel(result, "IsAuth"); value == false {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Unauthorized!")
	} else {
		tools.JsonResponse(result.Result, w)
	}
}
