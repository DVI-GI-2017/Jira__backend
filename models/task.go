package models

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Initiator   *User `json:"initiator"`
	Assignee    *User `json:"assignee"`
	Labels      Labels `json:"labels"`
}

type Tasks []User
