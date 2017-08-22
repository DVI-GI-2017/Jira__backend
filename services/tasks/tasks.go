package tasks

import (
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "tasks"

func CheckExistence(mongo *mgo.Database, task *models.Task) (bool, error) {
	c, err := mongo.C(collection).Find(bson.M{"title": task.Title}).Count()
	return c != 0, err
}

func Create(mongo *mgo.Database, task interface{}) (interface{}, error) {
	err := mongo.C(collection).Insert(task)
	if err != nil {
		return nil, err
	}
	err = mongo.C(collection).Find(task).One(task)
	return task, err
}

func All(mongo *mgo.Database) (result models.TasksList, err error) {
	const defaultSize = 100
	result = make(models.TasksList, defaultSize)

	err = mongo.C(collection).Find(bson.M{}).All(&result)
	return
}

func FindById(mongo *mgo.Database, id bson.ObjectId) (*models.Task, error) {
	task := models.NewTask()
	err := mongo.C(collection).FindId(id).One(task)
	return task, err
}
