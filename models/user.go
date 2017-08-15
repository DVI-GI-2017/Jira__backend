package models

type User struct {
	Name string `json:"name"`
	Data string `json:"Data"`
	Phone string `json:"Phone"`
}

type Users []User
