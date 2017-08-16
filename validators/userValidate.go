package validators

import (
	"errors"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/db"
)

func CheckUser(user models.User) (err error) {
	for _, value := range db.FakeUsers {
		if user.Email == value.Email && user.Password == value.Password {
			return nil
		}
	}

	return errors.New("Invalid credentials")
}
