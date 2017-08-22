package tasks

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const collection = "tasks"

func CheckExistence(source db.DataSource, task *models.Task) (bool, error) {
	c, err := source.C(collection).Find(bson.M{"title": task.Title}).Count()
	return c != 0, err
}

func Create(source db.DataSource, task interface{}) (interface{}, error) {
	err := source.C(collection).Insert(task)
	if err != nil {
		return nil, err
	}
	err = source.C(collection).Find(task).One(task)
	return task, err
}

func All(source db.DataSource) (result models.TasksList, err error) {
	result = make(models.TasksList, 0)

	err = source.C(collection).Find(bson.M{}).All(&result)
	return
}

func FindById(source db.DataSource, id bson.ObjectId) (*models.Task, error) {
	task := models.NewTask()
	err := source.C(collection).FindId(id).One(task)
	return task, err
}
