package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
	"gopkg.in/mgo.v2/bson"
)

func init() {
	resolvers["Task"] = tasksResolver
}

const (
	TaskCreate   = Action("TaskCreate")
	TaskExists   = Action("TaskExists")
	TasksAll     = Action("TasksAll")
	TaskFindById = Action("TaskFindById")
)

func tasksResolver(action Action) (service ServiceFunc) {
	switch action {

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

	default:
		log.Panicf("can not find resolver with action: %v, in tasks resolvers", action)
		return
	}
}
