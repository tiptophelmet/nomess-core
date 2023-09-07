package postprocessor

import (
	"net/http"
)

type PostProcFunc func(http.ResponseWriter, *http.Request)

var postProc = make(map[string][]PostProcFunc)

func Register(pattern string, postProcFuncList []PostProcFunc) {
	postProc[pattern] = postProcFuncList
}

func WithPostProcessor(w http.ResponseWriter, r *http.Request) {
	postProcList, found := postProc[r.URL.Path]
	if !found {
		return
	}

	for _, proc := range postProcList {
		proc(w, r)
	}
}
