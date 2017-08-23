package handlers

import (
	"fmt"
	"net/http"

	"encoding/json"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"gopkg.in/mgo.v2/bson"
)

// Creates project
// Post body - project
// Returns created project if OK
func CreateProject(w http.ResponseWriter, req *http.Request) {
	var projectInfo models.Project

	body := params.ExtractParams(req).Body

	err := json.Unmarshal(body, &projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := projectInfo.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.DispatchAction(pool.CheckProjectExists, projectInfo)

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

	project, err := pool.DispatchAction(pool.CreateProject, projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadGateway)
		return
	}

	JsonResponse(w, project)
}

// Returns all projects
func AllProjects(w http.ResponseWriter, _ *http.Request) {
	projects, err := pool.DispatchAction(pool.AllProjects, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, projects.(models.ProjectsList))
}

// Returns project with given id
// Query param: "id" - project id
func GetProjectById(w http.ResponseWriter, req *http.Request) {
	id := params.ExtractParams(req).PathParams["id"]

	user, err := pool.DispatchAction(pool.FindProjectById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user.(models.Project))
	return
}
