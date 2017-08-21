package handlers

import (
	"fmt"
	"net/http"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

func AllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := pool.DispatchAction(pool.AllUsers, nil)
	if err != nil {
		fmt.Fprint("Can not return all users!")
		log.Printf("Can not return all users: %v", err)
	}

	tools.JsonResponse(users.(models.UsersList), w)
}
