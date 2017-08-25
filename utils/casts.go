package utils

import "fmt"

// Helper to generate error message about bad cast.
func ErrInvalidCast(got, expected interface{}) error {
	return fmt.Errorf("can not cast input data with type: %T to %T", got, expected)
}
