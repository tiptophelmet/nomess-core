package responder

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v2/postprocessor"
)

type Responder struct {
	w http.ResponseWriter
	r *http.Request
}

var resp *Responder

func Init(w http.ResponseWriter, r *http.Request) *Responder {
	if resp != nil {
		return resp
	}

	resp = &Responder{w, r}
	return resp
}

func Respond(response interface{}, statusCode int) {
	w, _ := postprocessor.WithPostProcessor(resp.w, resp.r)

	resp.w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
