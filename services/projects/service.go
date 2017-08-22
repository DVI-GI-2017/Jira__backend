package projects

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "projects"

func CheckExistence(source db.DataSource, project *models.Project) (bool, error) {
	c, err := source.C(collection).Find(bson.M{"title": project.Title}).Count()
	return c != 0, err
}

func Create(source db.DataSource, project interface{}) (result interface{}, err error) {
	return project, source.C(collection).Insert(project)
}

func All(source db.DataSource) (result models.ProjectsList, err error) {
	const defaultSize = 100
	result = make(models.ProjectsList, defaultSize)

	err = source.C(collection).Find(bson.M{}).All(&result)
	return
}

func FindById(mongo *mgo.Database, id bson.ObjectId) (*models.Project, error) {
	project := new(models.Project)
	err := mongo.C(collection).FindId(id).One(project)
	return project, err
}
