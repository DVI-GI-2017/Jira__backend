package services

import (
	"github.com/DVI-GI-2017/Jira__backend/configs"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"gopkg.in/mgo.v2/bson"
)

func GetUserByEmailAndPassword(email string, password string) (result models.User, err error) {
	users := db.Connection.GetCollection(configs.ConfigInfo.Mongo)

	result = models.User{}
	err = users.Find(bson.M{
		"$and": []interface{}{
			bson.M{"email": email},
			bson.M{"password": password},
		},
	}).One(&result)

	return
}
