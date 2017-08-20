package handlers

import (
	"fmt"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

func AllUsers(w http.ResponseWriter, _ *http.Request) {
	var users models.UsersList

	action, _ := pool.NewAction(pool.AllUsers)

	pool.Queue <- &pool.Job{
		ModelType: users,
		Action:    action,
	}

	result := <-pool.Results

	if value := tools.GetValueFromModel(result, "IsAuth"); value == false {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprintln(w, "Unauthorized!")
	} else {
		tools.JsonResponse(result.ResultType, w)
	}
}
