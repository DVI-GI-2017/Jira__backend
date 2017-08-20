package tools

import (
	"errors"
	"reflect"
	"strconv"
)

func GetType(modelType interface{}) string {
	if t := reflect.TypeOf(modelType); t.Kind() == reflect.Ptr {
		return t.Elem().Name()
	} else {
		return t.Name()
	}
}

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

func SetParam2Model(model interface{}, key string, param interface{}) error {
	field := reflect.ValueOf(&model).Elem().FieldByName(key)

	switch field.Kind() {
	case reflect.Int:
		field.SetInt(param.(int64))
		return nil
	case reflect.Bool:
		field.SetBool(param.(bool))
		return nil
	case reflect.String:
		field.SetString(param.(string))
		return nil
	default:
		return errors.New("Can't set param in model!")
	}
}
