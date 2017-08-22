package projects

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const collection = "projects"

func CheckExistence(source db.DataSource, project models.Project) (bool, error) {
	empty, err := source.C(collection).Find(bson.M{"title": project.Title}).IsEmpty()
	return !empty, err
}

func Create(source db.DataSource, project models.Project) (result interface{}, err error) {
	return source.C(collection).Insert(project)
}

func All(source db.DataSource) (models.ProjectsList, error) {
	result, err := source.C(collection).Find(bson.M{}).All()
	if err != nil {
		return models.ProjectsList{}, err
	}
	return result.(models.ProjectsList), nil
}

func FindById(mongo db.DataSource, id bson.ObjectId) (models.Project, error) {
	result, err := mongo.C(collection).FindId(id).One()
	if err != nil {
		return models.Project{}, err
	}
	return result.(models.Project), nil
}
