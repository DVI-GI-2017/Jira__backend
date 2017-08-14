package models

type Project struct {
	id          uint32
	title       string
	description string
	tasks       []Task
}
