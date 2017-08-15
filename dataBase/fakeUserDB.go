package dataBase

import "Jira__backend/models"

var UsersListFromFakeDB = models.Users{
	models.User{Name: "User1", Data: "21.08.1997", Phone: "8(999)999-99-99"},
	models.User{Name: "User2", Data: "10.01.1997", Phone: "8(999)999-99-99"},
}
