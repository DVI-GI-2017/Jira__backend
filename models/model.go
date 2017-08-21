package models

const (
	UserType      = "User"
	UsersListType = "UsersList"

	TaskType      = "Task"
	TasksListType = "TasksList"

	ProjectType      = "Project"
	ProjectsListType = "ProjectsList"

	LabelType      = "Label"
	LabelsListType = "LabelsList"
)

func GetModel(modelName string) interface{} {
	switch modelName {
	case UserType:
		return new(User)
	case UsersListType:
		return new(UsersList)

	case TaskType:
		return new(Task)
	case TasksListType:
		return new(TasksList)

	case LabelType:
		return new(Label)
	case LabelsListType:
		return new(LabelsList)

	case ProjectType:
		return new(Project)
	case ProjectsListType:
		return new(ProjectsList)

	default:
		return nil
	}
}
