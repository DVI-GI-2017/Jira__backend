package pool

import "errors"

var typesActionList = [...]string{
	"Insert", "Find", "Insert and Find",
}
var handlersKeys = [...]string{
	"GetUser",
}

type FinderHandler func(model interface{}) interface{}

type Action struct {
	Type    string
	Finders map[string]FinderHandler
}

func InitFinderHnadlers() {
	for _, value := range handlersKeys{
		switch value {
		case "GetUser":
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
