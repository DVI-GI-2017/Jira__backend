package models

type Project struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Tasks       *Tasks `json:"tasks"`
}

type Projects []Project
