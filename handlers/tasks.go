package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"gopkg.in/mgo.v2/bson"
)

// Create task
// Post body - task
// Returns created task if OK
func CreateTask(w http.ResponseWriter, req *http.Request) {
	body := params.ExtractParams(req).Body

	var task models.Task

	err := json.Unmarshal(body, &task)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	exists, err := pool.DispatchAction(pool.CheckTaskExists, task)
	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("Task with title: %s already exists!", task.Title), http.StatusConflict)
		return
	}

	newTask, err := pool.DispatchAction(pool.CreateTask, task)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadGateway)
		return
	}

	JsonResponse(w, newTask)
}

// Returns all tasks
func AllTasks(w http.ResponseWriter, _ *http.Request) {
	tasks, err := pool.DispatchAction(pool.AllTasks, nil)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, tasks.(models.TasksList))
}

// Returns task with given id
// Path params: "id" - task id.
func GetTaskById(w http.ResponseWriter, req *http.Request) {

	id := params.ExtractParams(req).PathParams["id"]

	task, err := pool.DispatchAction(pool.FindTaskById, bson.ObjectIdHex(id))
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, task.(models.Task))
	return
}
