package validators

import (
	"errors"
	"Jira__backend/models"
	"Jira__backend/dataBase"
)

func CheckUser(user models.User) (err error) {
	for _, value := range dataBase.UsersListFromFakeDB {
		if user.Email == value.Email && user.Password == value.Password {
			return nil
		}
	}

	return errors.New("Invalid credentials")
}
