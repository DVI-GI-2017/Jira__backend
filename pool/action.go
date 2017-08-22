package pool

import (
	"errors"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
	"gopkg.in/mgo.v2/bson"
)

const (
	//Users actions
	CreateUser           = "CreateUser"
	CheckUserExists      = "CheckUserExists"
	CheckUserCredentials = "CheckUserCredentials"
	FindUserById         = "FindUserById"
	AllUsers             = "AllUsers"

	// Projects actions
	CreateProject      = "CreateProject"
	CheckProjectExists = "CheckProjectExists"
	AllProjects        = "AllProjects"
	FindProjectById    = "FindProjectById"

	// Tasks actions
	CreateTask      = "CreateTask"
	CheckTaskExists = "CheckTaskExists"
	AllTasks        = "AllTasks"
	FindTaskById    = "FindTaskById"

	// Labels actions
	AddLabelToTask       = "AddLabelToTask"
	AllLabelsOnTask      = "AllLabelsOnTask"
	CheckLabelAlreadySet = "CheckLabelAlreadySet"
)

var typesActionList = [...]string{
	// Users actions
	CreateUser,
	CheckUserExists,
	CheckUserCredentials,
	FindUserById,
	AllUsers,

	// Projects actions
	CreateProject,
	CheckProjectExists,
	AllProjects,
	FindProjectById,

	// Tasks actions
	CreateTask,
	CheckTaskExists,
	AllTasks,
	FindTaskById,

	// Labels actions
	AddLabelToTask,
	CheckLabelAlreadySet,
	AllLabelsOnTask,
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

type ServiceFunc func(source db.DataSource, data interface{}) (result interface{}, err error)

func GetServiceByAction(action *Action) (service ServiceFunc, err error) {
	switch action.Type {

	case CreateUser:
		service = func(source db.DataSource, user interface{}) (result interface{}, err error) {
			return users.CreateUser(source, user.(models.User))
		}
		return

	case CheckUserExists:
		service = func(source db.DataSource, credentials interface{}) (result interface{}, err error) {
			return users.CheckExistence(source, credentials.(models.User))
		}
		return

	case CheckUserCredentials:
		service = func(source db.DataSource, credentials interface{}) (interface{}, error) {
			return users.CheckCredentials(source, credentials.(models.User))
		}
		return

	case AllUsers:
		service = func(source db.DataSource, _ interface{}) (result interface{}, err error) {
			return users.AllUsers(source)
		}
		return

	case FindUserById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return users.FindUserById(source, id.(bson.ObjectId))
		}
		return

	case CreateProject:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.Create(source, project.(models.Project))
		}
		return

	case CheckProjectExists:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CheckExistence(source, project.(models.Project))
		}
		return

	case AllProjects:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.All(source)
		}
		return

	case FindProjectById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return projects.FindById(source, id.(bson.ObjectId))
		}
		return

	case CreateTask:
		service = func(source db.DataSource, task interface{}) (interface{}, error) {
			return tasks.CreateTask(source, task.(models.Task))
		}
		return

	case CheckTaskExists:
		service = func(source db.DataSource, task interface{}) (interface{}, error) {
			return tasks.CheckTaskExists(source, task.(models.Task))
		}
		return

	case AllTasks:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return tasks.AllTasks(source)
		}
		return

	case FindTaskById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return tasks.FindById(source, id.(bson.ObjectId))
		}
		return

	case AddLabelToTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			dataList := data.([]interface{})

			id := dataList[0].(bson.ObjectId)
			label := dataList[1].(models.Label)

			return tasks.AddLabelToTask(source, id, label)
		}
		return

	case CheckLabelAlreadySet:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			dataList := data.([]interface{})

			id := dataList[0].(bson.ObjectId)
			label := dataList[1].(models.Label)

			return tasks.CheckLabelAlreadySet(source, id, label)
		}
		return

	case AllLabelsOnTask:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return tasks.AllLabels(source, id.(bson.ObjectId))
		}
		return
	}

	return NullHandler, errors.New("Can't find handler!")
}

// Helper handler for case when handler not found.
func NullHandler(_ db.DataSource, _ interface{}) (result interface{}, err error) {
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
