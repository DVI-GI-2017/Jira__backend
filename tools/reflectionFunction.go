package tools

import (
	"errors"
	"reflect"
	"strconv"
)

func GetValueFromModel(model interface{}, key string) (interface{}, bool) {
	s := reflect.ValueOf(model)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	if s.Kind() != reflect.Struct {
		return "", false
	}
	f := s.FieldByName(key)
	if !f.IsValid() {
		return "", false
	}
	switch f.Kind() {
	case reflect.String:
		return f.Interface().(string), true
	case reflect.Bool:
		return f.Interface().(bool), true
	case reflect.Int:
		return strconv.FormatInt(f.Int(), 10), true
	default:
		return "", false
	}
}

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
	s := reflect.ValueOf(model)
	if s.Kind() == reflect.Ptr {
		s = s.Elem()
	}
	if s.Kind() != reflect.Struct {
		return errors.New("Not struct")
	}
	f := s.FieldByName(key)
	if !f.IsValid() {
		return errors.New("Not valid")
	}

	switch f.Kind() {
	case reflect.String:
		f.SetString(param.(string))
		return nil
	case reflect.Bool:
		f.SetBool(param.(bool))
		return nil
	case reflect.Int:
		f.SetInt(param.(int64))
		return nil
	default:
		return nil
	}
}
