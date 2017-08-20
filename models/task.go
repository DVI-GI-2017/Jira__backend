package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Initiator   *User     `json:"initiator"`
	Assignee    *User     `json:"assignee"`
	Labels      LabelsList   `json:"labels"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TasksList []Task