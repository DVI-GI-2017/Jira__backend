package yac

import (
	"log"
	"net/http"
	"time"
)

type WrapperFunc func(handlerFunc http.HandlerFunc) http.HandlerFunc

// Wraps handler func with slice of wrapper functions one by one.
func Wrap(h http.HandlerFunc, wrappers ...WrapperFunc) http.HandlerFunc {
	for _, w := range wrappers {
		h = w(h)
	}
	return h
}

// Logs requests
func Logger(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer log.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)

		h(w, r)
	}
}
