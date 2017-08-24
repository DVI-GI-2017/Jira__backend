package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"gopkg.in/mgo.v2/bson"
)

// Creates project
// Post body - project
// Returns created project if OK
func CreateProject(w http.ResponseWriter, req *http.Request) {
	var projectInfo models.Project

	body := mux.Params(req).Body

	err := json.Unmarshal(body, &projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := projectInfo.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.Dispatch(pool.ProjectExists, projectInfo)

	if err != nil {
		JsonErrorResponse(w, fmt.Errorf("can not check project existence: %v", err),
			http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("project with title %s already exists", projectInfo.Title),
			http.StatusConflict)
		return
	}

	project, err := pool.Dispatch(pool.ProjectCreate, projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadGateway)
		return
	}

	JsonResponse(w, project)
}

// Returns all projects
func AllProjects(w http.ResponseWriter, _ *http.Request) {
	projects, err := pool.Dispatch(pool.ProjectsAll, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, projects.(models.ProjectsList))
}

// Returns project with given id
// Query param: "id" - project id
func GetProjectById(w http.ResponseWriter, req *http.Request) {
	id := mux.Params(req).PathParams["id"]

	user, err := pool.Dispatch(pool.ProjectFindById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user.(models.Project))
	return
}

func GetAllUsersFromProject(w http.ResponseWriter, req *http.Request) {
	id := mux.Params(req).PathParams["id"]

	users, err := pool.Dispatch(pool.ProjectFindById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users.(models.UsersList))
}

func AddUserToProject(w http.ResponseWriter, req *http.Request) {
	projectId := mux.Params(req).PathParams["id"]
	userId := string(mux.Params(req).Body)

	users, err := pool.Dispatch(pool.ProjectAddUser,
		models.ProjectUser{
			ProjectId: models.RequiredId(bson.ObjectIdHex(projectId)),
			UserId:    models.RequiredId(bson.ObjectIdHex(userId)),
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users.(models.UsersList))
}

func DeleteUserFromProject(w http.ResponseWriter, req *http.Request) {
	projectId := mux.Params(req).PathParams["id"]
	userId := string(mux.Params(req).Body)

	users, err := pool.Dispatch(pool.ProjectDeleteUser,
		models.ProjectUser{
			ProjectId: models.RequiredId(bson.ObjectIdHex(projectId)),
			UserId:    models.RequiredId(bson.ObjectIdHex(userId)),
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users.(models.UsersList))
}
