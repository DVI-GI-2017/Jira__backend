package models

import (
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description,omitempty"`
	InitiatorId bson.ObjectId `json:"initiator_id" bson:"initiator_id,omitempty"`
	AssigneeId  bson.ObjectId `json:"assignee_id" bson:"assignee_id,omitempty"`
	Labels      LabelsList    `json:"labels" bson:"labels,omitempty"`
}

type TasksList []Task

// Initialize Labels with empty LabelsList.
func NewTask() *Task {
	return &Task{
		Labels: make(LabelsList, 0),
	}
}

// Safely copies task and return pointer to new task.
func (t *Task) Copy() *Task {
	newTask := &Task{
		Id:          t.Id,
		Title:       t.Title,
		Description: t.Description,
		InitiatorId: t.InitiatorId,
		AssigneeId:  t.AssigneeId,
		Labels:      t.Labels,
	}

	return newTask
}

func (t *Task) HasLabel(label Label) bool {
	for _, l := range t.Labels {
		if l == label {
			return true
		}
	}

	return false
}

func (t *Task) AddLabel(label Label) {
	t.Labels = append(t.Labels, label)
}
