package tasks

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const cTasks = "tasks"

// Checks if task with this 'title == task.Title' exists.
func CheckTaskExists(source db.DataSource, task models.Task) (bool, error) {
	c, err := source.C(cTasks).Find(bson.M{"title": task.Title}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check if task exists: %v", err)
	}
	return !c, nil
}

// Creates task and returns it.
func CreateTask(source db.DataSource, task models.Task) (models.Task, error) {
	task.Id = bson.NewObjectId()

	err := source.C(cTasks).Insert(task)
	if err != nil {
		return models.Task{}, fmt.Errorf("can not create task '%v': %v", task, err)
	}
	return task, nil
}

// Returns all tasks.
func AllTasks(source db.DataSource) (tasksList models.TasksList, err error) {
	err = queryAllTasks(source).All(&tasksList)
	if err != nil {
		return models.TasksList{}, fmt.Errorf("can not retrieve all tasks: %v", err)
	}
	return tasksList, nil
}

// Returns query with all tasks
func queryAllTasks(source db.DataSource) db.Query {
	return source.C(cTasks).Find(bson.M{})
}

// Returns task with given id
func FindTaskById(source db.DataSource, id bson.ObjectId) (task models.Task, err error) {
	err = queryTaskById(source, id).One(&task)
	if err != nil {
		return models.Task{}, fmt.Errorf("can not find task with id '%s': %s",
			id.Hex(), err)
	}
	return task, nil
}

// Returns query with possibly one task with matching id
func queryTaskById(source db.DataSource, id bson.ObjectId) db.Query {
	return source.C(cTasks).FindId(id)
}

// Updates task and returns new if OK.
func UpdateTask(source db.DataSource, task_id bson.ObjectId, update interface{}) (models.Task, error) {
	task, err := FindTaskById(source, task_id)
	if err != nil {
		return models.Task{},
			fmt.Errorf("can not find task with id '%s': %s", task_id.Hex(), err)
	}

	err = source.C(cTasks).Update(bson.M{"_id": task.Id}, update)
	if err != nil {
		return models.Task{}, fmt.Errorf("can not update task: %v", err)
	}

	task, err = FindTaskById(source, task_id)
	if err != nil {
		return models.Task{},
			fmt.Errorf("can not find updated task with id '%s': %s", task_id.Hex(), err)
	}

	return task, err
}
