package tasks

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

// Checks if label already set on this task.
func CheckLabelAlreadySet(source db.DataSource, id bson.ObjectId, label models.Label) (bool, error) {
	task, err := FindTaskById(source, id)

	if err != nil {
		return false, err
	}

	return task.HasLabel(label), nil
}

func AllLabels(source db.DataSource, id bson.ObjectId) (models.LabelsList, error) {
	task, err := FindTaskById(source, id)
	if err != nil {
		return models.LabelsList{}, err
	}

	return task.Labels, nil
}

// Adds label to task and returns new list of labels on this task.
func AddLabelToTask(source db.DataSource, task_id bson.ObjectId, label models.Label) (models.LabelsList, error) {
	task, err := FindTaskById(source, task_id)
	if err != nil {
		return models.LabelsList{},
			fmt.Errorf("can not find task with id '%s' :%v", label, task_id, err)
	}

	labels := append(task.Labels, label)

	return labels, source.C(cTasks).Update(task, bson.M{"labels": labels})
}
