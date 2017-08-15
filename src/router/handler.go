package router

import (
	"net/http"
	"encoding/json"
	"common/response"
)

type Response struct {
	Code  int
	Body  interface{}
	Error error
}

type handler func(*http.Request) Response

type handlerWriter func(http.ResponseWriter, *http.Request)

func (h handlerWriter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h(w, r)
}

type HTTPStatusReporter interface {
	HTTPStatusCode() int
}

func NewResponse(body interface{}, err error) (re Response) {
	re.Body, re.Error = body, err
	if err == nil {
		re.Code = http.StatusOK
		return
	}
	if err, test := err.(HTTPStatusReporter); test {
		re.Code = err.HTTPStatusCode()
	} else {
		re.Code = http.StatusInternalServerError
	}
	return
}

type HTTPError struct {
	Code    int
	Message string
}

func (e HTTPError) Error() string {
	return e.Message
}

func (e HTTPError) HTTPStatusCode() int {
	return e.Code
}