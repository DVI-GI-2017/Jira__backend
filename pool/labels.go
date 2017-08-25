package pool

import (
	"log"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
)

func init() {
	resolvers["Label"] = labelsResolver
}

const (
	LabelAddToTask      = Action("LabelAddToTask")
	LabelsAllOnTask     = Action("LabelsAllOnTask")
	LabelAlreadySet     = Action("LabelAlreadySet")
	LabelDeleteFromTask = Action("LabelDeleteFromTask")
)

func labelsResolver(action Action) ServiceFunc {
	switch action {

	case LabelAddToTask:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if taskLabel, ok := data.(models.TaskLabel); ok {
				return tasks.AddLabelToTask(source, taskLabel.TaskId, taskLabel.Label)
			}
			return models.LabelsList{}, castFailsMsg(data, models.TaskLabel{})
		}

	case LabelsAllOnTask:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return tasks.AllLabels(source, id)
			}
			return models.LabelsList{}, castFailsMsg(data, models.RequiredId{})
		}

	case LabelAlreadySet:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if taskLabel, ok := data.(models.TaskLabel); ok {
				return tasks.CheckLabelAlreadySet(source, taskLabel.TaskId, taskLabel.Label)
			}
			return false, castFailsMsg(data, models.TaskLabel{})
		}

	case LabelDeleteFromTask:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.DeleteLabelFromTask(source, taskLabel.TaskId, taskLabel.Label)
		}

	default:
		log.Panicf("can not find resolver with action: %v, in labels resolvers", action)
		return nil
	}
}
