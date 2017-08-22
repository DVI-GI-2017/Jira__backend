package users

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

const cUsers = "users"

// Checks if user with this credentials.Email exists.
func CheckUserExists(source db.DataSource, credentials models.User) (bool, error) {
	empty, err := source.C(cUsers).Find(bson.M{"email": credentials.Email}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check if user with credentials '%v' exists: %v", credentials, err)
	}
	return !empty, nil
}

// Checks if user credentials present in users collection.
func CheckUserCredentials(source db.DataSource, credentials models.User) (bool, error) {
	empty, err := source.C(cUsers).Find(credentials).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check user credentials '%v': %v", credentials, err)
	}
	return !empty, nil
}

// Creates user and returns it.
func CreateUser(source db.DataSource, user models.User) (models.User, error) {
	result, err := source.C(cUsers).Insert(user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not create user '%v': %v", user, err)
	}
	return result.(models.User), nil
}

// Returns all users.
func AllUsers(source db.DataSource) (models.UsersList, error) {
	result, err := source.C(cUsers).Find(nil).All()
	if err != nil {
		return models.UsersList{}, fmt.Errorf("can not retrieve all users: %v", err)
	}
	return result.(models.UsersList), nil
}

// Returns user with given id.
func FindUserById(source db.DataSource, id bson.ObjectId) (models.User, error) {
	user, err := source.C(cUsers).FindId(id).One()
	if err != nil {
		return models.User{}, fmt.Errorf("can not find user with id '%s': %v", id, err)
	}
	return user.(models.User), nil
}
