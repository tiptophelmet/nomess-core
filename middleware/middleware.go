package middleware

import (
	"net/http"
)

type MiddlewareFunc func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)

var defaultMwList []MiddlewareFunc

var mw = make(map[string][]MiddlewareFunc)

func Register(pattern string, mwFuncList []MiddlewareFunc) {
	mw[pattern] = mwFuncList
}

func Default(mwFuncList []MiddlewareFunc) {
	defaultMwList = mwFuncList
}

func WithMiddleware(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	mwList := defaultMwList

	if mwPathList, found := mw[r.URL.Path]; found {
		mwList = append(mwList, mwPathList...)
	}

	for _, mw := range mwList {
		w, r = mw(w, r)
	}

	return w, r
}
