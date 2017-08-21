package models

import "gopkg.in/mgo.v2/bson"

type Credentials struct {
	Id       bson.ObjectId `json:"_id" bson:"_id"`
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}
