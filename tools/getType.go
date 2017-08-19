package tools

import "reflect"

func getType(modelType interface{}) string {
	if t := reflect.TypeOf(modelType); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}
