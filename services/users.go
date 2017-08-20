package services

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
)

func Insert(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Insert("users", user)
}

func GetUserByEmailAndPassword(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Find("users", user)
}

func NullHandler(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return nil, nil
}
