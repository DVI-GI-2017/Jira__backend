package pool

import "errors"

var typesList = [...]string{
	"Insert", "Find", "Insert and Find",
}

type Action struct {
	Type string
}

func NewAction(actionType string) (*Action, error) {
	if checkType(actionType) {
		return &Action{
			Type: actionType,
		}, nil
	} else {
		return &Action{}, errors.New("Can't create new action!")
	}
}

func checkType(actionType string) bool {
	for _, value := range typesList {
		if value == actionType {
			return true
		}
	}

	return false
}
