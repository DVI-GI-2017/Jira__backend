package users

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const collection = "users"

func CheckExistence(source db.DataSource, credentials *models.User) (bool, error) {
	c, err := source.C(collection).Find(bson.M{"email": credentials.Email}).Count()
	return c != 0, err
}

func CheckCredentials(source db.DataSource, credentials *models.User) (bool, error) {
	c, err := source.C(collection).Find(credentials).Count()
	return c != 0, err
}

func Insert(source db.DataSource, user interface{}) (result interface{}, err error) {
	return user, source.C(collection).Insert(user)
}

func All(source db.DataSource) (result models.UsersList, err error) {
	const defaultSize = 100
	result = make(models.UsersList, defaultSize)

	err = source.C(collection).Find(nil).All(&result)
	return
}

func FindUserById(source db.DataSource, id bson.ObjectId) (*models.User, error) {
	user := new(models.User)
	err := source.C(collection).FindId(id).One(user)
	return user, err
}
