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

	case UserCreate:
		service = func(source db.DataSource, user interface{}) (result interface{}, err error) {
			return users.CreateUser(source, user.(models.User))
		}
		return

	case UserExists:
		service = func(source db.DataSource, credentials interface{}) (result interface{}, err error) {
			return users.CheckUserExists(source, credentials.(models.User))
		}
		return

	case UserAuthorize:
		service = func(source db.DataSource, credentials interface{}) (interface{}, error) {
			return users.CheckUserCredentials(source, credentials.(models.User))
		}
		return

	case UsersAll:
		service = func(source db.DataSource, _ interface{}) (result interface{}, err error) {
			return users.AllUsers(source)
		}
		return

	case UserFindById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return users.FindUserById(source, id.(bson.ObjectId))
		}
		return

	case ProjectCreate:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CreateProject(source, project.(models.Project))
		}
		return

	case ProjectExists:
		service = func(source db.DataSource, project interface{}) (interface{}, error) {
			return projects.CheckProjectExists(source, project.(models.Project))
		}
		return

	case ProjectsAll:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return projects.AllProjects(source)
		}
		return

	case ProjectFindById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return projects.FindProjectById(source, id.(bson.ObjectId))
		}
		return

	case TaskCreate:
		service = func(source db.DataSource, task interface{}) (interface{}, error) {
			return tasks.CreateTask(source, task.(models.Task))
		}
		return

	case TaskExists:
		service = func(source db.DataSource, task interface{}) (interface{}, error) {
			return tasks.CheckTaskExists(source, task.(models.Task))
		}
		return

	case TasksAll:
		service = func(source db.DataSource, _ interface{}) (interface{}, error) {
			return tasks.AllTasks(source)
		}
		return

	case TaskFindById:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return tasks.FindTaskById(source, id.(bson.ObjectId))
		}
		return

	case LabelAddToTask:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.AddLabelToTask(source, taskLabel.TaskId, taskLabel.Label)
		}
		return

	case LabelsAllOnTask:
		service = func(source db.DataSource, id interface{}) (interface{}, error) {
			return tasks.AllLabels(source, id.(bson.ObjectId))
		}
		return

	case LabelAlreadySet:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.CheckLabelAlreadySet(source, taskLabel.TaskId, taskLabel.Label)
		}
		return
	case LabelDeleteFromTask:
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
