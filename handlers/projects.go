package handlers

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/tools"
	"gopkg.in/mgo.v2/bson"
	"net/http"
)

var Projects = func(w http.ResponseWriter, request *http.Request) {
	result := models.Projects{}
	project := db.Connection.GetCollection(configs.ConfigInfo.Mongo.Db, db.ProjectCollection)
	err := project.Find(nil).All(&result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Project with id not found!")
		return
	} else {
		tools.JsonResponse(result, w)
	}
	tools.JsonResponse(db.FakeProjects, w)
}

var Project = func(w http.ResponseWriter, request *http.Request) {
	if !bson.IsObjectIdHex(request.URL.Query().Get(":id")) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Invalid id format!")
		return
	}
	result := models.Project{}
	project := db.Connection.GetCollection(configs.ConfigInfo.Mongo.Db, db.ProjectCollection)

	err := project.FindId(bson.ObjectIdHex(request.URL.Query().Get(":id"))).One(&result)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Project with id not found!")
		return
	} else {
		tools.JsonResponse(result, w)
	}
}
