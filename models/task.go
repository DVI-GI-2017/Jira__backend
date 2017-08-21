package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description,omitempty"`
	Initiator   *User         `json:"initiator" bson:"initiator,omitempty"`
	Assignee    *User         `json:"assignee" bson:"assignee,omitempty"`
	Labels      LabelsList    `json:"labels" bson:"labels,omitempty"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt   time.Time     `json:"updated_at" bson:"updated_at,omitempty"`
}

type TasksList []Task
