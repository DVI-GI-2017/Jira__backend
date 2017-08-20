package users

import "github.com/DVI-GI-2017/Jira__backend/db"

func Insert(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Insert("users", user)
}

func All(mongo *db.MongoConnection, _ interface{}) (result interface{}, err error) {
	return mongo.FindAll("users")
}

func GetUserByEmailAndPassword(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Find("users", user)
}
