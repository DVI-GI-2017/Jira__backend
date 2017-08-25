package models

import "fmt"

// Helper to generate error message about bad cast.
func ErrInvalidCast(got, expected interface{}) error {
	return fmt.Errorf("can not cast input data with type: %T to %T", got, expected)
}

func SafeCastToTaskLabel(data interface{}) (TaskLabel, error) {
	if val, ok := data.(TaskLabel); ok {
		return val, ErrInvalidCast(data, TaskLabel{})
	}
	return TaskLabel{}, nil
}

func SafeCastToRequiredId(data interface{}) (RequiredId, error) {
	if val, ok := data.(RequiredId); ok {
		return val, ErrInvalidCast(data, RequiredId{})
	}
	return RequiredId{}, nil
}

func SafeCastToProject(data interface{}) (Project, error) {
	if val, ok := data.(Project); ok {
		return val, ErrInvalidCast(data, Project{})
	}
	return Project{}, nil
}

func SafeCastToProjectUser(data interface{}) (ProjectUser, error) {
	if val, ok := data.(ProjectUser); ok {
		return val, ErrInvalidCast(data, ProjectUser{})
	}
	return ProjectUser{}, nil
}
