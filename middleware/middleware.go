package middleware

import (
	"net/http"
)

type MiddlewareFunc func(http.ResponseWriter, *http.Request)

var mw = make(map[string][]MiddlewareFunc)

func Register(pattern string, mwFuncList []MiddlewareFunc) {
	mw[pattern] = mwFuncList
}

func WithMiddleware(w http.ResponseWriter, r *http.Request) {
	mwList, found := mw[r.URL.Path]
	if !found {
		return
	}

	for _, mw := range mwList {
		mw(w, r)
	}
}
