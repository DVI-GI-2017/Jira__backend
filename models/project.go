package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description"`
	Tasks       TasksList     `json:"tasks" bson:"tasks,omitempty"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
}

type ProjectsList []Project
