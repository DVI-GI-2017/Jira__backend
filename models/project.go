package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Project struct {
	Id          bson.ObjectId
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Tasks       TasksList `json:"tasks"`
	CreatedAt   time.Time   `json:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at"`
}

type ProjectsList []Project