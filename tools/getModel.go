package tools

import "Jira__backend/models"

const (
	User        = "User"
	Task        = "Task"
	Project     = "Project"
	Credentials = "Credentials"
)

func GetModel(modelName string) interface{} {
	switch modelName {
	case User:
		return new(models.User)
	case Credentials:
		return new(models.User)
	case Task:
		return new(models.Task)
	case Project:
		return new(models.Project)
	default:
		return nil
	}
}
