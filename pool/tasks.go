package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
	"github.com/DVI-GI-2017/Jira__backend/utils"
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

func tasksResolver(action Action) (ServiceFunc, error) {
	switch action {

	case TaskCreate:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if task, ok := data.(models.Task); ok {
				return tasks.AddTaskToProject(source, task)
			}
			return models.Task{}, utils.ErrCastFails(data, models.Task{})
		}, nil

	case TaskExists:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if task, ok := data.(models.Task); ok {
				return tasks.CheckTaskExists(source, task)
			}
			return models.Task{}, utils.ErrCastFails(data, models.Task{})
		}, nil

	case TasksAllOnProject:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return tasks.AllTasks(source, id)
			}
			return models.TasksList{}, utils.ErrCastFails(data, models.RequiredId{})
		}, nil

	case TaskFindById:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return tasks.FindTaskById(source, id)
			}
			return models.Task{}, utils.ErrCastFails(data, models.RequiredId{})
		}, nil

	default:
		return nil, fmt.Errorf("can not find resolver with action: %v, in tasks resolvers", action)
	}
}
