package models

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Assignee    *User `json:"assignee"`
	Labels      Labels `json:"labels"`
}

type Tasks []User
