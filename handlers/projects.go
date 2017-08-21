package handlers

import (
	"fmt"
	"log"
	"net/http"

	"encoding/json"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/params"
	"github.com/DVI-GI-2017/Jira__backend/pool"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"gopkg.in/mgo.v2/bson"
)

func CreateProject(w http.ResponseWriter, req *http.Request) {
	body := params.ExtractParams(req).Body

	projectInfo := new(models.Project)

	err := json.Unmarshal(body, projectInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	exists, err := pool.DispatchAction(pool.CheckProjectExists, projectInfo)
	if exists.(bool) {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Project with title: %s already exists!", projectInfo.Title)

		log.Printf("Project with title: %s already exists!", projectInfo.Title)

		return
	}

	project, err := pool.DispatchAction(pool.CreateProject, projectInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "Can not create project. Please, try later")
		log.Printf("can not create project: %v", err)

		return
	}

	tools.JsonResponse(project, w)
}

func AllProjects(w http.ResponseWriter, _ *http.Request) {
	projects, err := pool.DispatchAction(pool.AllProjects, nil)
	if err != nil {
		fmt.Fprint(w, "Can not return all projects!")
		log.Printf("Can not return all projects: %v", err)

		return
	}

	tools.JsonResponse(projects.(models.ProjectsList), w)
}

func GetProjectById(w http.ResponseWriter, req *http.Request) {
	parameters := params.ExtractParams(req).PathParams

	if id, ok := parameters["id"]; ok {
		user, err := pool.DispatchAction(pool.FindProjectById, bson.ObjectIdHex(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Printf("Can not find task by id: %v because of: %v", id, err)
			return
		}

		tools.JsonResponse(user.(*models.Project), w)
		return
	}

	http.NotFound(w, req)
}
