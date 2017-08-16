package auth

import (
	"github.com/DVI-GI-2017/Jira__backend/db"
	"errors"
)

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func LoginUser(user *Credentials) (err error) {
	for _, value := range db.FakeUsers {
		if user.Email == value.Email && user.Password == value.Password {
			// TODO: Login user
			return nil
		}
	}

	return errors.New("Invalid credentials")
}
