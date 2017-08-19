package tools

import (
	"reflect"
	"strconv"
)

type ParseModelMap map[string]string

func ParseModel(model interface{}) ParseModelMap {
	parseMap := make(map[string]string)
	val := reflect.ValueOf(model).Elem()

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)

		parseMap[typeField.Name] = valueField.Interface()
	}

	return parseMap
}

func interface2String(object interface{}) {
	if GetType(object) == "int" {
		strconv.Itoa(object)
	}
}
