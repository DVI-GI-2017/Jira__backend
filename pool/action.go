package pool

import (
	"errors"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

const (
	Insert        = "Insert"
	Find          = "Find"
	InsertAndFind = "Insert and Find"
)

var typesActionList = [...]string{
	Insert, Find, InsertAndFind,
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
	case Insert:
		return services.Insert, nil
	case Find:
		return services.GetUserByEmailAndPassword, nil
	case InsertAndFind:
		break
	}

	return services.NullHandler, errors.New("Can't find handler!")
}
