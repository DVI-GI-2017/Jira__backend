package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

func AddLabelToTask(w http.ResponseWriter, req *http.Request) {
	parameters := params.ExtractParams(req)

	label := new(models.Label)
	err := json.Unmarshal(parameters.Body, label)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	id := parameters.PathParams["id"]

	exists, err := pool.DispatchAction(pool.CheckLabelAlreadySet, []interface{}{id, label})
	if err != nil {
		log.Printf("can not check label presence: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if exists.(bool) {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprint(w, "label already set")
		log.Printf("label %s already set on task %s", label, id)
		return
	}

	_, err = pool.DispatchAction(pool.AddLabelToTask, []interface{}{id, label})
	if err != nil {
		log.Printf("can not add label to task: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func AllLabelsOnTask(w http.ResponseWriter, req *http.Request) {
	pathParams := params.ExtractParams(req).PathParams
	id := pathParams["id"]

	labels, err := pool.DispatchAction(pool.AllLabelsOnTask, id)
	if err != nil {
		fmt.Fprint(w, "Can not return all labels!")
		log.Printf("Can not return all labels: %v", err)

		return
	}

	tools.JsonResponse(labels.(models.LabelsList), w)
}
