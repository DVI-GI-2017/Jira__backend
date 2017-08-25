package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
)

func init() {
	resolvers["Task"] = tasksResolver
}

const (
	TaskCreate        = Action("TaskCreate")
	TaskExists        = Action("TaskExists")
	TasksAllOnProject = Action("TasksAllOnProject")
	TaskFindById      = Action("TaskFindById")
)

func tasksResolver(action Action) ServiceFunc {
	switch action {

	case TaskCreate:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if task, ok := data.(models.Task); ok {
				return tasks.AddTaskToProject(source, task)
			}
			return models.Task{}, castFailsMsg(data, models.Task{})
		}

	case TaskExists:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if task, ok := data.(models.Task); ok {
				return tasks.CheckTaskExists(source, task)
			}
			return models.Task{}, castFailsMsg(data, models.Task{})
		}

	case TasksAllOnProject:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return tasks.AllTasks(source, id)
			}
			return models.TasksList{}, castFailsMsg(data, models.RequiredId{})
		}

	case TaskFindById:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return tasks.FindTaskById(source, id)
			}
			return models.Task{}, castFailsMsg(data, models.RequiredId{})
		}

	default:
		log.Panicf("can not find resolver with action: %v, in tasks resolvers", action)
		return nil
	}
}
