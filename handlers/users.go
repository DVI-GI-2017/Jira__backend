package handlers

import (
	"fmt"
	"net/http"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"gopkg.in/mgo.v2/bson"
)

func AllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := pool.DispatchAction(pool.AllUsers, nil)
	if err != nil {
		fmt.Fprint(w, "Can not return all users!")
		log.Printf("Can not return all users: %v", err)
	}

	tools.JsonResponse(users.(models.UsersList), w)
}

func GetUserById(w http.ResponseWriter, req *http.Request) {
	parameters := params.ExtractParams(req).PathParams

	if id, ok := parameters["id"]; ok {
		user, err := pool.DispatchAction(pool.FindUserById, bson.ObjectIdHex(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("Can not find user by id: %v", id)
			return
		}

		tools.JsonResponse(user.(*models.User), w)
		return
	}

	http.NotFound(w, req)
}
