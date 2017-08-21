package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId
	Title       string     `json:"title" bson:"title"`
	Description string     `json:"description" bson:"description"`
	Initiator   *User      `json:"initiator" bson:"initiator"`
	Assignee    *User      `json:"assignee" bson:"assignee"`
	Labels      LabelsList `json:"labels" bson:"labels"`
	CreatedAt   time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at" bson:"updated_at"`
}

type TasksList []Task
