package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	ProjectId   bson.ObjectId `json:"project_id" bson:"project_id"`
	Description string        `json:"description" bson:"description,omitempty"`
	InitiatorId bson.ObjectId `json:"initiator_id" bson:"initiator_id,omitempty"`
	AssigneeId  bson.ObjectId `json:"assignee_id" bson:"assignee_id,omitempty"`
	Labels      LabelsList    `json:"labels" bson:"labels,omitempty"`
}

type TasksList []Task

// Helper structure for label associated with task.
type TaskLabel struct {
	TaskId bson.ObjectId
	Label  Label
}
