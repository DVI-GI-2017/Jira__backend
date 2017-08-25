package models

import "fmt"

// Helper to generate error message about bad cast.
func ErrInvalidCast(got, expected interface{}) error {
	return fmt.Errorf("can not cast input data with type: %T to %T", got, expected)
}

// Helper to generate cast to 'RequiredId' error
func ErrInvalidCastToRequiredId(data interface{}) error {
	return ErrInvalidCast(data, RequiredId{})
}

// Helper to generate cast to 'TaskLabel' error
func ErrInvalidCastToTaskLabel(data interface{}) error {
	return ErrInvalidCast(data, TaskLabel{})
}

// Helper to generate cast to 'bool' error
func ErrInvalidCastToBool(data interface{}) error {
	return ErrInvalidCast(data, false)
}
