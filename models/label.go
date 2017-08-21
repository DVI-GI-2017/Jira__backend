package models

import "gopkg.in/mgo.v2/bson"

type Label struct {
	Id   bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

type LabelsList []Label
