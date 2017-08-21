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

func CreateTask(w http.ResponseWriter, req *http.Request) {
	body := params.ExtractParams(req).Body

	taskInfo := new(models.Task)

	fmt.Println(body)
	fmt.Println(taskInfo)
	err := json.Unmarshal(body, &taskInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprint(w, "Error in request!")
		log.Printf("%v", err)

		return
	}

	exists, err := pool.DispatchAction(pool.CheckTaskExists, taskInfo)
	if exists.(bool) {
		w.WriteHeader(http.StatusConflict)
		fmt.Fprintf(w, "Task with title: %s already exists!", taskInfo.Title)

		log.Printf("Task with title: %s already exists!", taskInfo.Title)

		return
	}

	project, err := pool.DispatchAction(pool.CreateTask, taskInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		fmt.Fprint(w, "Can not create task. Please, try later")
		log.Printf("can not create task: %v", err)

		return
	}

	tools.JsonResponse(project, w)
}

func AllTasks(w http.ResponseWriter, _ *http.Request) {
	projects, err := pool.DispatchAction(pool.AllTasks, nil)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		fmt.Fprint(w, "Can not return all tasks!")
		log.Printf("Can not return all tasks: %v", err)

		return
	}

	tools.JsonResponse(projects.(models.TasksList), w)
}

func GetTaskById(w http.ResponseWriter, req *http.Request) {
	parameters := params.ExtractParams(req).PathParams

	if id, ok := parameters["id"]; ok {
		task, err := pool.DispatchAction(pool.FindTaskById, bson.ObjectIdHex(id))
		if err != nil {
			w.WriteHeader(http.StatusNotFound)

			fmt.Fprintln(w, "Can't find task!")
			log.Printf("Can not find task by id: %v because of: %v", id, err)
			return
		}

		tools.JsonResponse(task.(*models.Task), w)
		return
	}

	http.NotFound(w, req)
}
