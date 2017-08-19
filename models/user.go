package models

import (
	"encoding/json"
	"fmt"
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
}

type Users []User

func (u *User) CopyMethod(user interface{}) {
	jsonModel, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("parse error: %s", err)

		return
	}

	json.Unmarshal(jsonModel, u)
}
