package tasks

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

func CheckLabelAlreadySet(source db.DataSource, id bson.ObjectId, label models.Label) (bool, error) {
	task, err := FindById(source, id)

	if err != nil {
		return false, err
	}

	return task.HasLabel(label), nil
}

func AllLabels(source db.DataSource, id bson.ObjectId) (*models.LabelsList, error) {
	task, err := FindById(source, id)
	if err != nil {
		return nil, err
	}

	return &task.Labels, nil
}

func AddLabelToTask(source db.DataSource, id bson.ObjectId, label models.Label) error {
	task, err := FindById(source, id)
	if err != nil {
		return err
	}

	newTask := task.Copy()
	newTask.AddLabel(label)

	return source.C(collection).Update(task, newTask)
}
