package models

import (
	"errors"

	"gopkg.in/mgo.v2/bson"
)

var ErrTaskNotFound = errors.New("Task not found")

type Project struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Description string        `json:"description" bson:"description,omitempty"`
	Tasks       TasksList     `json:"tasks" bson:"tasks,omitempty"`
}

type ProjectsList []Project

// Initialize Tasks with empty TasksList.
func NewProject() *Project {
	return &Project{
		Tasks: make(TasksList, 0),
	}
}

func (p *Project) HasTask(title string) bool {
	for _, task := range p.Tasks {
		if task.Title == title {
			return true
		}
	}

	return false
}

func (p *Project) FindTask(title string) (*Task, error) {
	for _, task := range p.Tasks {
		if task.Title == title {
			return &task, nil
		}
	}

	return nil, ErrTaskNotFound
}
