package handlers

import (
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"gopkg.in/mgo.v2/bson"
)

// Returns all users.
func AllUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := pool.DispatchAction(pool.AllUsers, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users.(models.UsersList))
}

// Returns user with given id.
// Path param: "id" - user id.
func GetUserById(w http.ResponseWriter, req *http.Request) {
	id := mux.Params(req).PathParams["id"]
	user, err := pool.DispatchAction(pool.FindUserById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user.(models.User))
}
