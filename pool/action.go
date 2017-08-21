package pool

import (
	"errors"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
	"gopkg.in/mgo.v2"
)

const (
	InsertUser           = "InsertUser"
	CheckUserExists      = "CheckUserExists"
	CheckUserCredentials = "CheckUserCredentials"
	FindUser             = "FindUser"
	AllUsers             = "AllUsers"
	UpdateUser           = "UpdateUser"
)

var typesActionList = [...]string{
	InsertUser, FindUser, CheckUserExists,
	CheckUserCredentials, FindUser, AllUsers,
	UpdateUser,
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

type ServiceFunc func(*mgo.Database, interface{}) (interface{}, error)

func GetServiceByAction(action *Action) (ServiceFunc, error) {
	switch action.Type {
	case InsertUser:
		return users.Insert, nil
	case CheckUserExists:
		return func(mongo *mgo.Database, credentials interface{}) (interface{}, error) {
			return users.CheckExistence(mongo, credentials.(*models.Credentials))
		}, nil
	case CheckUserCredentials:
		return func(mongo *mgo.Database, credentials interface{}) (interface{}, error) {
			return users.CheckCredentials(mongo, credentials.(*models.Credentials))
		}, nil
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
func NullHandler(_ *mgo.Database, _ interface{}) (result interface{}, err error) {
	return nil, nil
}

// Creates job with given action and input and returns result.
func DispatchAction(actionType string, input interface{}) (result interface{}, err error) {
	action, err := NewAction(actionType)
	if err != nil {
		log.Panicf("invalid actionType type: %s", actionType)
	}

	Queue <- &Job{
		Input:  input,
		Action: action,
	}

	jobResult := <-Results

	result = jobResult.Result
	err = jobResult.Error

	return
}
