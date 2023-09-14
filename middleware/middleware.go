package middleware

import (
	"net/http"
)

type MiddlewareFunc func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)

var mw = make(map[string][]MiddlewareFunc)

func Register(pattern string, mwFuncList []MiddlewareFunc) {
	mw[pattern] = mwFuncList
}

func WithMiddleware(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	mwList, found := mw[r.URL.Path]
	if !found {
		return w, r
	}

	for _, mw := range mwList {
		w, r = mw(w, r)
	}

	return w, r
}
