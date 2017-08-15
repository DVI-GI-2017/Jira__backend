package models

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Assignee    *User `json:"assignee"`
}

type Tasks []User
