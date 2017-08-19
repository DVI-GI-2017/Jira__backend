package tools

import (
	"fmt"
	"reflect"
	"strconv"
)

type ParseModelMap map[string]string

func ParseModel(model interface{}) (values ParseModelMap) {
	values = make(ParseModelMap)

	iVal := reflect.ValueOf(model).Elem()
	typ := iVal.Type()

	for i := 0; i < iVal.NumField(); i++ {
		f := iVal.Field(i)
		var v string

		switch f.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(f.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(f.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(f.Float(), 'f', 4, 64)
		case []byte:
			v = string(f.Bytes())
		case string:
			v = f.String()
		}

		values[typ.Field(i).Name] = v
	}
	return
}
