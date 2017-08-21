package projects

import (
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "labels"

func CheckExistence(mongo *mgo.Database, label *models.Label) (bool, error) {
	c, err := mongo.C(collection).Find(bson.M{"name": label.Name}).Count()
	return c != 0, err
}

func Create(mongo *mgo.Database, label interface{}) (result interface{}, err error) {
	return label, mongo.C(collection).Insert(label)
}

func All(mongo *mgo.Database) (result models.LabelsList, err error) {
	const defaultSize = 100
	result = make(models.LabelsList, defaultSize)

	err = mongo.C(collection).Find(bson.M{}).All(&result)
	return
}

func FindById(mongo *mgo.Database, id bson.ObjectId) (*models.Label, error) {
	label := new(models.Label)
	err := mongo.C(collection).FindId(id).One(label)
	return label, err
}
