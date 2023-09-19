package middleware

import (
	"net/http"

	"github.com/tiptophelmet/nomess-core/v5/util"
)

type MiddlewareFunc func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)

var defaultMwList []MiddlewareFunc

var mw = make(map[string][]MiddlewareFunc)

func Register(pattern string, mwFuncList []MiddlewareFunc) {
	mw[pattern] = mwFuncList
}

func RegisterMulti(patterns []string, mwFuncList []MiddlewareFunc) {
	for _, pattern := range patterns {
		mw[pattern] = mwFuncList
	}
}

func Default(mwFuncList []MiddlewareFunc) {
	defaultMwList = mwFuncList
}

func WithMiddleware(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	mwList := defaultMwList

	pattern := util.GetRoutePattern(r)

	if mwPathList, found := mw[pattern]; found {
		mwList = append(mwList, mwPathList...)
	}

	for _, mw := range mwList {
		w, r = mw(w, r)
	}

	return w, r
}
