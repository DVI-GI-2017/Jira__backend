package users

import "github.com/DVI-GI-2017/Jira__backend/db"

const collection = "Users"

func Insert(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Insert(collection, user)
}

func All(mongo *db.MongoConnection, _ interface{}) (result interface{}, err error) {
	return mongo.FindAll(collection)
}

func GetUserByEmailAndPassword(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Find(collection, user)
}
