package users

import (
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const collection = "Users"

func CheckExistence(mongo *mgo.Database, credentials *models.Credentials) (bool, error) {
	c, err := mongo.C(collection).Find(bson.M{"Email": credentials.Email}).Count()
	return c != 0, err
}

func CheckCredentials(mongo *mgo.Database, credentials *models.Credentials) (bool, error) {
	c, err := mongo.C(collection).Find(credentials).Count()
	return c != 0, err
}

func Insert(mongo *mgo.Database, user interface{}) (result interface{}, err error) {
	return user, mongo.C(collection).Insert(user)
}

func All(mongo *mgo.Database) (result models.UsersList, err error) {
	const defaultSize = 100
	result = make(models.UsersList, defaultSize)

	err = mongo.C(collection).Find(bson.M{}).All(&result)
	return
}

func FindUserById(mongo *mgo.Database, id bson.ObjectId) (*models.User, error) {
	user := new(models.User)
	err := mongo.C(collection).FindId(id).One(user)
	return user, err
}
