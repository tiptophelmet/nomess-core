package responder

import (
	"encoding/json"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v2/logger"
	"github.com/tiptophelmet/nomess-core/v2/postprocessor"
)

type WithResponseFunc func(response interface{}, statusCode int)

func Respond(w http.ResponseWriter, r *http.Request) WithResponseFunc {
	w, _ = postprocessor.WithPostProcessor(w, r)

	return func(response interface{}, statusCode int) {
		w.WriteHeader(statusCode)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			logger.Error("failed to encode response '%v' for %v", response, r.RequestURI)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}
