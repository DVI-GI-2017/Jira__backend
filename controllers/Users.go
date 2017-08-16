package controllers

import (
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

func CheckUser(user *auth.Credentials) (result models.User, err error) {
	return services.GetUserByEmailAndPassword(user.Email, user.Password)
}

func AddUser(user *auth.Credentials) (err error) {
	return services.AddUser(user)
}
