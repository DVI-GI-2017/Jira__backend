package models

type User struct {
	Name string `json:"name"`
	Data string `json:"data"`
	Phone string `json:"phone"`
}

type Users []User
