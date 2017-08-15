package tools

type ValidatorFunc func(v interface{}, param string) error
