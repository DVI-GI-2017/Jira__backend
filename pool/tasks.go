package pool

import (
	"fmt"

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

func tasksResolver(action Action) (service ServiceFunc, err error) {
	switch action {

	case TaskCreate:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			task, err := models.SafeCastToTask(data)
			if err != nil {
				return models.Task{}, err
			}
			return tasks.AddTaskToProject(source, task)
		}
		return

	case TaskExists:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			task, err := models.SafeCastToTask(data)
			if err != nil {
				return models.Task{}, err
			}
			return tasks.CheckTaskExists(source, task)
		}
		return

	case TasksAllOnProject:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.TasksList{}, err
			}
			return tasks.AllTasks(source, id)
		}
		return

	case TaskFindById:
		service = func(source db.DataSource, data interface{}) (interface{}, error) {
			id, err := models.SafeCastToRequiredId(data)
			if err != nil {
				return models.Task{}, err
			}
			return tasks.FindTaskById(source, id)
		}
		return
	}
	return nil, fmt.Errorf("can not find resolver with action: %v, in tasks resolvers", action)

}
