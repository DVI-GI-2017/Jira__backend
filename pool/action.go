package pool

import (
	"errors"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

var typesActionList = [...]string{
	"Insert", "Find", "Insert and Find",
}
var handlersKeys = [...]string{
	"GetUser",
}

var findersHandler map[string]FinderHandler

type FinderHandler func(*db.MongoConnection, interface{}) (interface{}, error)

type Action struct {
	Type    string
	Finders map[string]FinderHandler
}

func InitFinderHnadlers() {
	for _, value := range handlersKeys {
		switch value {
		case "GetUser":
			findersHandler[value] = services.GetUserByEmailAndPassword
			break
		}
	}
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
