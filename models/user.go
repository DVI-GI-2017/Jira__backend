package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	FirstName string        `json:"first_name" bson:"_id,omitempty"`
	LastName  string        `json:"last_name" bson:"_id,omitempty"`
	Bio       string        `json:"bio" bson:"_id,omitempty"`
	Tasks     TasksList     `json:"tasks" bson:"_id,omitempty"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
}

type UsersList []User
