package users

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const collectionUsers = "users"

func CheckExistence(source db.DataSource, credentials models.User) (bool, error) {
	empty, err := source.C(collectionUsers).Find(bson.M{"email": credentials.Email}).IsEmpty()
	return !empty, err
}

func CheckCredentials(source db.DataSource, credentials models.User) (bool, error) {
	empty, err := source.C(collectionUsers).Find(credentials).IsEmpty()
	return !empty, err
}

func CreateUser(source db.DataSource, user models.User) (result interface{}, err error) {
	return source.C(collectionUsers).Insert(user)
}

func AllUsers(source db.DataSource) (models.UsersList, error) {
	result, err := source.C(collectionUsers).Find(nil).All()
	if err != nil {
		return models.UsersList{}, err
	}
	return result.(models.UsersList), err
}

func FindUserById(source db.DataSource, id bson.ObjectId) (models.User, error) {
	user, err := source.C(collectionUsers).FindId(id).One()
	if err != nil {
		return models.User{}, err
	}
	return user.(models.User), err
}
