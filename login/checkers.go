package login

import (
	"errors"
	"github.com/DVI-GI-2017/Jira__backend/db"
)

func CheckUser(user *Credentials) (err error) {
	for _, value := range db.FakeUsers {
		if user.Email == value.Email && user.Password == value.Password {
			return nil
		}
	}

	return errors.New("Invalid credentials")
}
