package router

import (
	"context"
	"fmt"
	"net/http"

	"github.com/tiptophelmet/nomess-core/v5/logger"
)

type routePatternContextKey string

const key routePatternContextKey = "NOMESS-ROUTE-PATTERN"

func markRoutePattern(pattern string, r *http.Request) *http.Request {
	ctx := context.WithValue(r.Context(), key, pattern)
	return r.WithContext(ctx)
}

func GetRoutePattern(r *http.Request) (string, error) {
	routePatternContextVal, toStrOk := r.Context().Value(key).(string)
	if !toStrOk {
		err := fmt.Errorf("route pattern was not marked for %s", r.URL.RequestURI())

		logger.Error(err.Error())
		return "", err
	}

	return routePatternContextVal, nil
}
