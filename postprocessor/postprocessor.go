package postprocessor

import (
	"net/http"
)

type PostProcFunc func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)

var defaultPostProcList []PostProcFunc

var postProc = make(map[string][]PostProcFunc)

func Register(pattern string, postProcFuncList []PostProcFunc) {
	postProc[pattern] = postProcFuncList
}

func Default(postProcFuncList []PostProcFunc) {
	defaultPostProcList = postProcFuncList
}

func WithPostProcessor(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	postProcList := defaultPostProcList

	if postProcPathList, found := postProc[r.URL.Path]; found {
		postProcList = append(postProcList, postProcPathList...)
	}

	for _, proc := range postProcList {
		w, r = proc(w, r)
	}

	return w, r
}
