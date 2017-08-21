package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id        bson.ObjectId `json:"_id" bson:"_id"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Bio       string        `json:"bio"`
	Tasks     TasksList     `json:"tasks"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	IsAuth    bool          `json:"is_auth"`
}

type UsersList []User
