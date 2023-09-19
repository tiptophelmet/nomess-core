package config

import (
	"reflect"
	"strconv"

	"github.com/tiptophelmet/nomess-core/v5/logger"
	"github.com/tiptophelmet/nomess-core/v5/util"
)

type configOptions struct {
	name       string
	rawVal     interface{}
	isFallback bool
}

func Get(name string) *configOptions {
	env, found := confList.list[name]
	if !found {
		return nil
	}

	options := &configOptions{name: name}
	if env.value != nil {
		options.rawVal = env.value
	} else {
		options.rawVal = env.fallback
		options.isFallback = true
	}

	return options
}

func (co *configOptions) Required() *configOptions {
	if util.IsEmpty(co.rawVal) {
		logger.Panic("config '%v' is required", co.name)
		return nil
	}

	return co
}

func (co *configOptions) Str() string {
	val, typeOk := co.rawVal.(string)
	if !typeOk {
		logger.Error("config '%v' does not assert to string (suggested: '%v')",
			co.name, reflect.TypeOf(co.rawVal))

		return ""
	}

	return val
}

func (co *configOptions) Bool() bool {
	if co.isFallback {
		boolVal, typeOk := co.rawVal.(bool)
		if !typeOk {
			logger.Error("config '%v' does not assert to bool (suggested: '%v')",
				co.name, reflect.TypeOf(co.rawVal))

			return false
		}

		return boolVal
	}

	boolVal, err := strconv.ParseBool(co.Str())
	if err != nil {
		logger.Error("config '%v' does not convert to bool", co.name)
		return false
	}

	return boolVal
}

func (co *configOptions) Int() int {
	if co.isFallback {
		int64Val := co.Int64()
		if int64Val == 0 {
			return 0
		}

		return int(int64Val)
	}

	int64Val, err := strconv.ParseInt(co.Str(), 10, 0)
	if err != nil {
		logger.Error("config '%v' does not convert to int", co.name)
		return 0
	}

	return int(int64Val)
}

func (co *configOptions) Int32() int32 {
	if co.isFallback {
		int64Val := co.Int64()
		if int64Val == 0 {
			return 0
		}

		return int32(int64Val)
	}

	int64Val, err := strconv.ParseInt(co.Str(), 10, 32)
	if err != nil {
		logger.Error("config '%v' does not convert to int32", co.name)
		return 0
	}

	return int32(int64Val)
}

func (co *configOptions) Int64() int64 {
	if co.isFallback {
		int64Val, typeOk := co.rawVal.(int64)
		if !typeOk {
			logger.Error("config '%v' does not assert to int64 (suggested: '%v')",
				co.name, reflect.TypeOf(co.rawVal))

			return 0
		}

		return int64Val
	}

	int64Val, err := strconv.ParseInt(co.Str(), 10, 64)
	if err != nil {
		logger.Error("config '%v' does not convert to int64", co.name)
		return 0
	}

	return int64Val
}

func (co *configOptions) Float64() float64 {
	if co.isFallback {
		val, typeOk := co.rawVal.(float64)
		if !typeOk {
			logger.Error("config '%v' does not assert to float64 (suggested: '%v')",
				co.name, reflect.TypeOf(co.rawVal))

			return 0.0
		}

		return val
	}

	float64Val, err := strconv.ParseFloat(co.Str(), 64)
	if err != nil {
		logger.Error("config '%v' does not convert to float64", co.name)
		return 0
	}

	return float64Val
}
