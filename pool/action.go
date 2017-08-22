package pool

import (
	"errors"

	"log"

	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
	"gopkg.in/mgo.v2"
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

type ServiceFunc func(*mgo.Database, interface{}) (interface{}, error)

func GetServiceByAction(action *Action) (ServiceFunc, error) {
	switch action.Type {

	case CreateUser:
		return users.CreateUser, nil

	case CheckUserExists:
		return func(mongo *mgo.Database, credentials interface{}) (interface{}, error) {
			return users.CheckExistence(mongo, credentials.(*models.User))
		}, nil

	case CheckUserCredentials:
		return func(mongo *mgo.Database, credentials interface{}) (interface{}, error) {
			return users.CheckCredentials(mongo, credentials.(*models.User))
		}, nil

	case AllUsers:
		return func(mongo *mgo.Database, _ interface{}) (interface{}, error) {
			return users.AllUsers(mongo)
		}, nil

	case FindUserById:
		return func(mongo *mgo.Database, id interface{}) (interface{}, error) {
			return users.FindUserById(mongo, id.(bson.ObjectId))
		}, nil

	case CreateProject:
		return func(mongo *mgo.Database, project interface{}) (interface{}, error) {
			return projects.Create(mongo, project)
		}, nil

	case CheckProjectExists:
		return func(mongo *mgo.Database, project interface{}) (interface{}, error) {
			return projects.CheckExistence(mongo, project.(*models.Project))
		}, nil

	case AllProjects:
		return func(mongo *mgo.Database, _ interface{}) (interface{}, error) {
			return projects.All(mongo)
		}, nil

	case FindProjectById:
		return func(mongo *mgo.Database, id interface{}) (interface{}, error) {
			return projects.FindById(mongo, id.(bson.ObjectId))
		}, nil

	case CreateTask:
		return func(mongo *mgo.Database, task interface{}) (interface{}, error) {
			return tasks.Create(mongo, task)
		}, nil
	case CheckTaskExists:
		return func(mongo *mgo.Database, task interface{}) (interface{}, error) {
			return tasks.CheckExistence(mongo, task.(*models.Task))
		}, nil
	case AllTasks:
		return func(mongo *mgo.Database, _ interface{}) (interface{}, error) {
			return tasks.All(mongo)
		}, nil
	case FindTaskById:
		return func(mongo *mgo.Database, id interface{}) (interface{}, error) {
			return tasks.FindById(mongo, id.(bson.ObjectId))
		}, nil

	case AddLabelToTask:
		return func(mongo *mgo.Database, data interface{}) (interface{}, error) {
			dataList := data.([]interface{})

			id := dataList[0].(bson.ObjectId)
			label := dataList[1].(models.Label)

			return nil, tasks.AddLabelToTask(mongo, id, label)
		}, nil
	case CheckLabelAlreadySet:
		return func(mongo *mgo.Database, data interface{}) (interface{}, error) {
			dataList := data.([]interface{})

			id := dataList[0].(bson.ObjectId)
			label := dataList[1].(models.Label)

			return tasks.CheckLabelAlreadySet(mongo, id, label)
		}, nil
	case AllLabelsOnTask:
		return func(mongo *mgo.Database, id interface{}) (interface{}, error) {
			return tasks.AllLabels(mongo, id.(bson.ObjectId))
		}, nil
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
