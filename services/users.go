package services

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
)

//func GetUserByEmailAndPassword1(email string, password string) (result models.User, err error) {
//	users := db.Connection.GetCollection(db.UserCollection)
//
//	result = models.User{}
//	err = users.Find(bson.M{
//		"$and": []interface{}{
//			bson.M{"email": email},
//			bson.M{"password": password},
//		},
//	}).One(&result)
//
//	return
//}
//
//func AddUser(user *auth.Credentials) (err error) {
//	users := db.Connection.GetCollection(db.UserCollection)
//	err = users.Insert(user)
//
//	return
//}

func GetUserByEmailAndPassword(mongo *db.MongoConnection, user interface{}) (result interface{}, err error) {
	return mongo.Find("users", user)
}
