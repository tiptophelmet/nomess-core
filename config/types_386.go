package config

import (
	"reflect"

	"github.com/tiptophelmet/nomess-core/v2/logger"
	"github.com/tiptophelmet/nomess-core/v2/util"
)

type configOptions struct {
	name   string
	rawVal interface{}
}

func Get(name string) *configOptions {
	options := &configOptions{
		name:   name,
		rawVal: raw(name),
	}

	return options
}

func raw(name string) interface{} {
	env, found := confList.list[name]
	if !found {
		return nil
	}

	if env.value != nil {
		return env.value
	}

	return env.fallback
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
	val, typeOk := co.rawVal.(bool)
	if !typeOk {
		logger.Error("config '%v' does not assert to bool (suggested: '%v')",
			co.name, reflect.TypeOf(co.rawVal))

		return false
	}

	return val
}

func (co *configOptions) Int() int {
	valInt32 := co.Int32()
	if valInt32 == 0 {
		return 0
	}

	return int(valInt32)
}

func (co *configOptions) Int32() int32 {
	val, typeOk := co.rawVal.(int32)
	if !typeOk {
		logger.Error("config '%v' does not assert to int32 (suggested: '%v')",
			co.name, reflect.TypeOf(co.rawVal))

		return 0
	}

	return val
}

func (co *configOptions) Int64() int64 {
	val, typeOk := co.rawVal.(int64)
	if !typeOk {
		logger.Error("config '%v' does not assert to int64 (suggested: '%v')",
			co.name, reflect.TypeOf(co.rawVal))

		return 0
	}

	return val
}

func (co *configOptions) Float64() float64 {
	val, typeOk := co.rawVal.(float64)
	if !typeOk {
		logger.Error("config '%v' does not assert to float64 (suggested: '%v')",
			co.name, reflect.TypeOf(co.rawVal))

		return 0.0
	}

	return val
}
