package pool

import (
	"errors"
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
