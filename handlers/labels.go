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

// Returns all labels from task
// Path parameter: "task_id" - task id.
func AllLabelsOnTask(w http.ResponseWriter, req *http.Request) {
	pathParams := params.ExtractParams(req).PathParams
	id := bson.ObjectIdHex(pathParams["task_id"])

	labels, err := pool.DispatchAction(pool.AllLabelsOnTask, id)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}

// Adds label to task.
// Query parameter: "task_id" - task id.
// Post body - label.
func AddLabelToTask(w http.ResponseWriter, req *http.Request) {
	p := params.ExtractParams(req)

	var label models.Label
	err := json.Unmarshal(p.Body, &label)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = label.Validate()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	taskId := bson.ObjectIdHex(p.PathParams["task_id"])
	taskLabel := models.TaskLabel{TaskId: taskId, Label: label}

	exists, err := pool.DispatchAction(pool.CheckLabelAlreadySet, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if exists.(bool) {
		JsonErrorResponse(w, fmt.Errorf("label '%v' already set on project '%s'", label, taskId.Hex()),
			http.StatusConflict)
		return
	}

	labels, err := pool.DispatchAction(pool.AddLabelToTask, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}

// Deletes label from task and returns new labels
// Path parameter: "task_id" - task id.
// Post body - label
func DeleteLabelFromTask(w http.ResponseWriter, req *http.Request) {
	p := params.ExtractParams(req)

	var label models.Label
	err := json.Unmarshal(p.Body, &label)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	taskId := bson.ObjectIdHex(p.PathParams["task_id"])

	taskLabel := models.TaskLabel{TaskId: taskId, Label: label}

	labels, err := pool.DispatchAction(pool.DeleteLabelFromTask, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}
