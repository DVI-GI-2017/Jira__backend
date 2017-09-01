package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/weitbelou/yac"
)

// Returns all labels from task
// Path parameter: "task_id" - task id.
func AllLabelsOnTask(w http.ResponseWriter, req *http.Request) {
	params, err := yac.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	id, err := models.NewRequiredId(params.PathParams["task_id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	labels, err := pool.Dispatch(pool.LabelsAllOnTask, id)
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
	params, err := yac.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var label models.Label
	err = json.Unmarshal(params.Body, &label)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = label.Validate()
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	taskId, err := models.NewRequiredId(params.PathParams["task_id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	taskLabel := models.TaskLabel{TaskId: taskId, Label: label}

	labels, err := pool.Dispatch(pool.LabelAddToTask, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	JsonResponse(w, labels)
}

// Deletes label from task and returns new labels
// Path parameter: "task_id" - task id.
// Post body - label
func DeleteLabelFromTask(w http.ResponseWriter, req *http.Request) {
	params, err := yac.Params(req)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	var label models.Label
	err = json.Unmarshal(params.Body, &label)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusBadRequest)
		return
	}

	taskId, err := models.NewRequiredId(params.PathParams["task_id"])
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	taskLabel := models.TaskLabel{TaskId: taskId, Label: label}

	labels, err := pool.Dispatch(pool.LabelDeleteFromTask, taskLabel)
	if err != nil {
		JsonErrorResponse(w, err, http.StatusNotFound)
		return
	}

	JsonResponse(w, labels)
}
