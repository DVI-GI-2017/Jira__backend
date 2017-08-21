package tasks

import (
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func CheckLabelAlreadySet(mongo *mgo.Database, id bson.ObjectId, label models.Label) (bool, error) {
	task, err := FindById(mongo, id)

	if err != nil {
		return false, err
	}

	if task.Labels == nil {
		return false, nil
	}

	for _, l := range task.Labels {
		if l == label {
			return true, nil
		}
	}

	return false, nil
}

func AddLabelToTask(mongo *mgo.Database, id bson.ObjectId, label models.Label) error {
	task, err := FindById(mongo, id)
	if err != nil {
		return err
	}

	if task.Labels == nil {
		task.Labels = make(models.LabelsList, 0)
	}

	task.Labels = append(task.Labels, label)

	return mongo.C(collection).Update(bson.M{"_id": task.Id}, task)
}

func AllLabels(mongo *mgo.Database, id bson.ObjectId) (*models.LabelsList, error) {
	task, err := FindById(mongo, id)
	if err != nil {
		return nil, err
	}

	return &task.Labels, nil
}
