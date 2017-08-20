package models

import (
	"gopkg.in/mgo.v2/bson"
	"time"
)

type User struct {
	ID        bson.ObjectId `bson:"_id,omitempty"`
	Email     string        `json:"email"`
	Password  string        `json:"password"`
	FirstName string        `json:"first_name"`
	LastName  string        `json:"last_name"`
	Bio       string        `json:"bio"`
	Tasks     Tasks         `json:"tasks"`
	CreatedAt time.Time     `json:"created_at"`
	UpdatedAt time.Time     `json:"updated_at"`
	IsAuth    bool          `json:"is_auth"`
}

type Users []User
