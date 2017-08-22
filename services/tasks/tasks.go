package tasks

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const collectionTasks = "tasks"

func CheckTaskExists(source db.DataSource, task models.Task) (bool, error) {
	c, err := source.C(collectionTasks).Find(bson.M{"title": task.Title}).IsEmpty()
	return !c, err
}

func CreateTask(source db.DataSource, task models.Task) (models.Task, error) {
	newTask, err := source.C(collectionTasks).Insert(task)
	if err != nil {
		return models.Task{}, err
	}
	return newTask.(models.Task), nil
}

func AllTasks(source db.DataSource) (models.TasksList, error) {
	result, err := source.C(collectionTasks).Find(nil).All()
	if err != nil {
		return models.TasksList{}, err
	}
	return result.(models.TasksList), nil
}

func FindById(source db.DataSource, id bson.ObjectId) (models.Task, error) {
	result, err := source.C(collectionTasks).FindId(id).One()
	if err != nil {
		return models.Task{}, err
	}
	return result.(models.Task), nil
}
