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
	"gopkg.in/mgo.v2/bson"
)

func CreateLabel(w http.ResponseWriter, req *http.Request) {
	body := params.ExtractParams(req).Body

	labelInfo := new(models.Label)

	err := json.Unmarshal(body, labelInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	exists, err := pool.DispatchAction(pool.CheckLabelExists, labelInfo)
	if exists.(bool) {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Label with name: %s already exists!", labelInfo.Name)

		log.Printf("Label with name: %s already exists!", labelInfo.Name)

		return
	}

	label, err := pool.DispatchAction(pool.CreateLabel, labelInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "Can not create label. Please, try later")
		log.Printf("can not create label: %v", err)

		return
	}

	tools.JsonResponse(label, w)
}

func AllLabels(w http.ResponseWriter, _ *http.Request) {
	labels, err := pool.DispatchAction(pool.AllLabels, nil)
	if err != nil {
		fmt.Fprint(w, "Can not return all labels!")
		log.Printf("Can not return all labels: %v", err)

		return
	}

	tools.JsonResponse(labels.(models.LabelsList), w)
}

func GetLabelById(w http.ResponseWriter, req *http.Request) {
	parameters := params.ExtractParams(req).PathParams

	if id, ok := parameters["id"]; ok {
		label, err := pool.DispatchAction(pool.FindLabelById, bson.ObjectIdHex(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("Can not find label by id: %v because of: %v", id, err)
			return
		}

		tools.JsonResponse(label.(*models.Label), w)
		return
	}

	http.NotFound(w, req)
}
