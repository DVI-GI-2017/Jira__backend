package models

import "gopkg.in/mgo.v2/bson"

type Label struct {
	Id bson.ObjectId
	Name string `json:"name"`
}

type Labels []Label
