package handlers

import (
	"net/http"

	"encoding/json"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/weitbelou/yacrouter"
)

// Creates project
// Post body - project
// Returns created project if OK
func CreateProject(w http.ResponseWriter, req *http.Request) {
	var projectInfo models.Project

	params, err := mux.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}
	body := params.Body

	err = json.Unmarshal(body, &projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := projectInfo.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	project, err := pool.Dispatch(pool.ProjectCreate, projectInfo)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusConflict)
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

	JsonResponse(w, projects)
}

// Returns project with given id
// Query param: "id" - project id
func GetProjectById(w http.ResponseWriter, req *http.Request) {
	params, err := mux.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := models.NewRequiredId(params.PathParams["id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	user, err := pool.Dispatch(pool.ProjectFindById, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user)
	return
}

func GetAllUsersFromProject(w http.ResponseWriter, req *http.Request) {
	params, err := mux.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := models.NewRequiredId(params.PathParams["id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	users, err := pool.Dispatch(pool.ProjectAllUsers, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users)
}

func GetAllTasksFromProject(w http.ResponseWriter, req *http.Request) {
	params, err := mux.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := models.NewRequiredId(params.PathParams["id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	tasks, err := pool.Dispatch(pool.ProjectAllTasks, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, tasks)
}

func AddUserToProject(w http.ResponseWriter, req *http.Request) {
	params, err := mux.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	projectId, err := models.NewRequiredId(params.PathParams["id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	var userId models.RequiredId
	err = json.Unmarshal(params.Body, &userId)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	users, err := pool.Dispatch(pool.ProjectAddUser,
		models.ProjectUser{
			ProjectId: projectId,
			UserId:    userId,
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, users)
}

func DeleteUserFromProject(w http.ResponseWriter, req *http.Request) {
	params, err := mux.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	projectId, err := models.NewRequiredId(params.PathParams["id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	var userId models.RequiredId
	err = json.Unmarshal(params.Body, &userId)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	user, err := pool.Dispatch(pool.ProjectDeleteUser,
		models.ProjectUser{
			ProjectId: projectId,
			UserId:    userId,
		})
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, user)
}
