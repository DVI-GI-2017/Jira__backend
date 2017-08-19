package services

import (
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

func GetUserByEmailAndPassword1(email string, password string) (result models.User, err error) {
	users := db.Connection.GetCollection(db.UserCollection)

	result = models.User{}
	err = users.Find(bson.M{
		"$and": []interface{}{
			bson.M{"email": email},
			bson.M{"password": password},
		},
	}).One(&result)

	return
}

func AddUser(user *auth.Credentials) (err error) {
	users := db.Connection.GetCollection(db.UserCollection)
	err = users.Insert(user)

	return
}

func GetUserByEmailAndPassword(mongo *db.MongoConnection, user auth.Credentials) (result models.User, err error) {
	users := mongo.GetCollection("users")

	result = models.User{}
	err = users.Find(bson.M{
		"$and": user,
	}).One(&result)

	return result, err
}
