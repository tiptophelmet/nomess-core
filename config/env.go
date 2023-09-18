package config

import (
	"os"

	"github.com/tiptophelmet/nomess-core/v4/logger"
)

func initEnv(name string, fallback interface{}) *env {
	var value interface{}

	if envValue, present := os.LookupEnv(name); present {
		value = envValue
	} else {
		value = nil
	}

	if value == nil && fallback == nil {
		logger.Fatal("env %s is not set & has no fallback", name)
	}

	return &env{name, value, fallback}
}

type env struct {
	name     string
	value    interface{}
	fallback interface{}
}
