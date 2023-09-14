package postprocessor

import (
	"net/http"
)

type PostProcFunc func(http.ResponseWriter, *http.Request) (http.ResponseWriter, *http.Request)

var postProc = make(map[string][]PostProcFunc)

func Register(pattern string, postProcFuncList []PostProcFunc) {
	postProc[pattern] = postProcFuncList
}

func WithPostProcessor(w http.ResponseWriter, r *http.Request) (http.ResponseWriter, *http.Request) {
	postProcList, found := postProc[r.URL.Path]
	if !found {
		return w, r
	}

	for _, proc := range postProcList {
		w, r = proc(w, r)
	}

	return w, r
}
