package router

import (
	"net/http"
)

func HttpStatus(inner handler) handler {
	return func(r *http.Request) (response Response) {
		response = inner(r)
		if response.Code == 0 {
			if response.Error == nil {
				response.Code = http.StatusOK
			} else {
				response.Code = http.StatusInternalServerError
			}
		}
		return
	}
}
