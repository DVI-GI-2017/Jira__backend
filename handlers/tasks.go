package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/mux"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"gopkg.in/mgo.v2/bson"
)

// Create task
// Post body - task
// Returns created task if OK
func CreateTask(w http.ResponseWriter, req *http.Request) {
	body := mux.Params(req).Body

	var task models.Task
	if err := json.Unmarshal(body, &task); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err := task.Validate(); err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.Dispatch(pool.CheckTaskExists, task)
	if err != nil {
		JsonErrorResponse(w, fmt.Errorf("can not check task existence: %v", err),
			http.StatusInternalServerError)
		return
	}
	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("Task with title: %s already exists!", task.Title),
			http.StatusConflict)
		return
	}

	newTask, err := pool.Dispatch(pool.CreateTask, task)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadGateway)
		return
	}

	JsonResponse(w, newTask)
}

// Returns all tasks
func AllTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := pool.Dispatch(pool.AllTasks, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, tasks.(models.TasksList))
}

// Returns task with given id
// Path params: "id" - task id.
func GetTaskById(w http.ResponseWriter, req *http.Request) {

	id := mux.Params(req).PathParams["id"]

	task, err := pool.Dispatch(pool.FindTaskById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, task.(models.Task))
	return
}
