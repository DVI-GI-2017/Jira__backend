package pool

import (
	"errors"
	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/services"
)

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
