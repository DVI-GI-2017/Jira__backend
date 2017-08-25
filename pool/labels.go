package pool

import (
	"fmt"

	"github.com/DVI-GI-2017/Jira__backend/db"
	"github.com/DVI-GI-2017/Jira__backend/models"
	"github.com/DVI-GI-2017/Jira__backend/services/tasks"
	"github.com/DVI-GI-2017/Jira__backend/utils"
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

func labelsResolver(action Action) (ServiceFunc, error) {
	switch action {

	case LabelAddToTask:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if taskLabel, ok := data.(models.TaskLabel); ok {
				return tasks.AddLabelToTask(source, taskLabel.TaskId, taskLabel.Label)
			}
			return models.LabelsList{}, utils.ErrInvalidCast(data, models.TaskLabel{})
		}, nil

	case LabelsAllOnTask:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if id, ok := data.(models.RequiredId); ok {
				return tasks.AllLabels(source, id)
			}
			return models.LabelsList{}, utils.ErrInvalidCast(data, models.RequiredId{})
		}, nil

	case LabelAlreadySet:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			if taskLabel, ok := data.(models.TaskLabel); ok {
				return tasks.CheckLabelAlreadySet(source, taskLabel.TaskId, taskLabel.Label)
			}
			return false, utils.ErrInvalidCast(data, models.TaskLabel{})
		}, nil

	case LabelDeleteFromTask:
		return func(source db.DataSource, data interface{}) (interface{}, error) {
			taskLabel := data.(models.TaskLabel)

			return tasks.DeleteLabelFromTask(source, taskLabel.TaskId, taskLabel.Label)
		}, nil

	default:
		return nil, fmt.Errorf("can not find resolver with action: %v, in labels resolvers", action)
	}
}
