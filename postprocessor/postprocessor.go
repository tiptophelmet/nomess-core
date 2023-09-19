package postprocessor

import (
	"net/http"

	"github.com/tiptophelmet/nomess-core/v5/util"
)

type PostProcFunc func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)

var defaultPostProcList []PostProcFunc

var postProc = make(map[string][]PostProcFunc)

func Register(pattern string, postProcFuncList []PostProcFunc) {
	postProc[pattern] = postProcFuncList
}

func RegisterMulti(patterns []string, postProcFuncList []PostProcFunc) {
	for _, pattern := range patterns {
		postProc[pattern] = postProcFuncList
	}
}

func Default(postProcFuncList []PostProcFunc) {
	defaultPostProcList = postProcFuncList
}

func WithPostProcessor(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	postProcList := defaultPostProcList

	pattern := util.GetRoutePattern(r)

	if postProcPathList, found := postProc[pattern]; found {
		postProcList = append(postProcList, postProcPathList...)
	}

	for _, proc := range postProcList {
		w, r = proc(w, r)
	}

	return w, r
}
