package tasks

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

// Returns all labels from given task.
func AllLabels(source db.DataSource, task_id bson.ObjectId) (models.LabelsList, error) {
	var container = struct {
		LabelsList models.LabelsList `bson:"labels"`
	}{}

	err := queryLabels(source.C(cTasks), task_id).One(&container)

	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not retrieve all labels on task %s: %v", task_id.Hex(), err)
	}

	return container.LabelsList, nil
}

// Selects labels from tasks query
func queryLabels(collection db.Collection, task_id bson.ObjectId) db.Query {
	return collection.Find(
		bson.M{
			"_id": task_id,
		}).Select(bson.M{"labels": 1})
}

// Checks if label already set on this task.
func CheckLabelAlreadySet(source db.DataSource, id bson.ObjectId, label models.Label) (bool, error) {
	notset, err := queryLabel(source.C(cTasks), id, label).IsEmpty()
	if err != nil {
		return false, err
	}

	return !notset, nil
}

// Selects label from collection.
func queryLabel(collection db.Collection, task_id bson.ObjectId, label models.Label) db.Query {
	return collection.Find(
		bson.M{
			"_id":    task_id,
			"labels": label,
		}).Select(bson.M{"labels": 1})
}

// Adds label to task and returns new list of labels on this task.
func AddLabelToTask(source db.DataSource, task_id bson.ObjectId, label models.Label) (models.LabelsList, error) {
	task, err := UpdateTask(source, task_id, pushLabel(label))
	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not add label '%v' to task '%s': %v", label, task.Id.Hex(), err)
	}

	return task.Labels, nil
}

// Returns "updater" that performs push in task labels array.
func pushLabel(label models.Label) interface{} {
	return bson.M{"$push": bson.M{"labels": label}}
}

// Deletes label from task and returns new list of labels on this task
func DeleteLabelFromTask(source db.DataSource, task_id bson.ObjectId, label models.Label) (models.LabelsList, error) {
	task, err := UpdateTask(source, task_id, bson.M{"$pull": bson.M{"labels": label}})
	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not delete label '%v' from task '%s': %v", label, task.Id.Hex(), err)
	}

	return task.Labels, nil
}
