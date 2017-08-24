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
	empty, err := source.C(cUsers).Find(bson.M{
		"email":    credentials.Email,
		"password": credentials.Password,
	}).IsEmpty()
	if err != nil {
		return false, fmt.Errorf("can not check user credentials '%v': %v", credentials, err)
	}
	return !empty, nil
}

// Creates user and returns it.
func CreateUser(source db.DataSource, user models.User) (models.User, error) {
	user.Id = models.AutoId(bson.NewObjectId())

	err := source.C(cUsers).Insert(user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not create user '%v': %v", user, err)
	}
	return user, nil
}

// Returns all users.
func AllUsers(source db.DataSource) (usersLists models.UsersList, err error) {
	err = source.C(cUsers).Find(nil).All(&usersLists)
	if err != nil {
		return models.UsersList{}, fmt.Errorf("can not retrieve all users: %v", err)
	}
	return usersLists, nil
}

// Returns user with given id.
func FindUserById(source db.DataSource, id bson.ObjectId) (user models.User, err error) {
	err = source.C(cUsers).FindId(id).One(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not find user with id '%s': %v", id, err)
	}
	return user, nil
}

// Returns user with given email.
func FindUserByEmail(source db.DataSource, email models.Email) (user models.User, err error) {
	err = source.C(cUsers).Find(bson.M{"email": email}).Select(bson.M{"text": 1}).One(&user)
	if err != nil {
		return models.User{}, fmt.Errorf("can not find user with email '%s': %v", email, err)
	}
	return user, nil
}
