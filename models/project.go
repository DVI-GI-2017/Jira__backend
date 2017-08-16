package models

import "gopkg.in/mgo.v2/bson"

type Project struct {
	Id          bson.ObjectId
	Title       string `json:"title"`
	Description string `json:"description"`
	Tasks       Tasks `json:"tasks"`
}

type Projects []Project
