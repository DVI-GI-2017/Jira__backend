package services

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/configs"
)

func GetUserByEmailAndPassword(email string, password string) (result models.User, err error) {
	connection := db.NewDBConnection(configs.ConfigInfo.Mongo)
	defer connection.CloseConnection()

	users := connection.GetCollection(configs.ConfigInfo.Mongo)

	result = models.User{}
	err = users.Find(bson.M{
		"$and": []interface{}{
			bson.M{"email": email},
			bson.M{"password": password},
		},
	}).One(&result)

	return
}
