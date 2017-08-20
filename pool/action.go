package pool

import (
	"errors"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
)

const (
	InsertUser = "InsertUser"
	FindUser   = "FindUser"
	AllUsers   = "AllUsers"
	UpdateUser = "UpdateUser"
)

var typesActionList = [...]string{
	InsertUser, FindUser, UpdateUser,
}

type Action struct {
	Type string
}

func NewAction(actionType string) (*Action, error) {
	if checkActionType(actionType) {
		return &Action{
			Type: actionType,
		}, nil
	} else {
		return &Action{}, errors.New("Can't create new action!")
	}
}

func checkActionType(actionType string) bool {
	for _, value := range typesActionList {
		if value == actionType {
			return true
		}
	}

	return false
}

type ServiceFunc func(*db.MongoConnection, interface{}) (interface{}, error)

func GetServiceByAction(action *Action) (ServiceFunc, error) {
	switch action.Type {
	case InsertUser:
		return users.Insert, nil
	case FindUser:
		return users.GetUserByEmailAndPassword, nil
	case AllUsers:
		return users.All, nil
	case UpdateUser:
		break
	}

	return NullHandler, errors.New("Can't find handler!")
}

// Helper handler for case when handler not found.
func NullHandler(_ *db.MongoConnection, _ interface{}) (result interface{}, err error) {
	return nil, nil
}
