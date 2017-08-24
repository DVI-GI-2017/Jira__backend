package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/projects"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
	"github.com/DVI-GI-2017/Jira__backend/services/users"
	"gopkg.in/mgo.v2/bson"
)

type ServiceFunc func(source db.DataSource, data interface{}) (result interface{}, err error)

func getServiceByAction(action Action) (service ServiceFunc, err error) {
	switch action {

	case CreateUser:
		service = func(source db.DataSource, user interface{}) (result interface{}, err error) {
			return users.CreateUser(source, user.(models.User))
		}
		return

	case CheckUserExists:
		service = func(source db.DataSource, credentials interface{}) (result interface{}, err error) {
			return users.CheckUserExists(source, credentials.(models.User))
		}
		return

	case CheckUserCredentials:
		service = func(source db.DataSource, credentials interface{}) (interface{}, error) {
			return users.CheckUserCredentials(source, credentials.(models.User))
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
			return projects.CreateProject(source, project.(models.Project))
		}
		return

	case CheckProjectExists:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CheckProjectExists(source, project.(models.Project))
		}
		return

	case AllProjects:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.AllProjects(source)
		}
		return

	case FindProjectById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return projects.FindProjectById(source, id.(bson.ObjectId))
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
			return tasks.FindTaskById(source, id.(bson.ObjectId))
		}
		return

	case AddLabelToTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.AddLabelToTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return

	case AllLabelsOnTask:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return tasks.AllLabels(source, id.(bson.ObjectId))
		}
		return

	case CheckLabelAlreadySet:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.CheckLabelAlreadySet(source, taskLabel.TaskId, taskLabel.Label)
		}
		return
	case DeleteLabelFromTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.DeleteLabelFromTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return
	default:
		log.Panicf("unknown action: %s", action)
		return
	}
}

// Creates job with given action and input and returns result.
func Dispatch(action Action, input interface{}) (result interface{}, err error) {
	Queue <- &Job{
		Input:  input,
		Action: action,
	}

	jobResult := <-Results

	result = jobResult.Result
	err = jobResult.Error

	return
}
