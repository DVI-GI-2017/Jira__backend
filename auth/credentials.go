package auth

import (
	"errors"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/configs"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(user *Credentials) (err error) {
	for _, value := range db.FakeUsers {
		if user.Email == value.Email && user.Password == value.Password {
			// TODO: Login user
			return nil
		}
	}

	return errors.New("Invalid credentials")
}

func RegisterUser(user *Credentials) (result models.User, err error) {
	connection := db.NewDBConnection(configs.ConfigInfo.Mongo)
	defer connection.CloseConnection()

	users := connection.GetCollection(configs.ConfigInfo.Mongo)

	result = models.User{}
	err = users.Find(bson.M{
		"$and": []interface{}{
			bson.M{"email": user.Email},
			bson.M{"password": user.Password},
		},
	}).One(&result)

	fmt.Println("Result")
	fmt.Println(result)

	return
}
