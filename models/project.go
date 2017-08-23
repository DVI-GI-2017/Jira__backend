package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description,omitempty"`
}

type ProjectsList []Project
