package config

import (
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

	var rawVal interface{}

	if env.value != nil {
		rawVal = env.value
	} else {
		rawVal = env.fallback
	}

	return rawVal
}

func (co *configOptions) Required() *configOptions {
	if util.IsEmpty(co.rawVal) {
		logger.Panic("could not resolve config %v", co.name)
		return nil
	}

	return co
}

func (co *configOptions) Str() string {
	val, typeOk := co.rawVal.(string)
	if !typeOk {
		logger.Error("could not assert config %v to string", co.name)
		return ""
	}

	return val
}

func (co *configOptions) Int() int {
	val, typeOk := co.rawVal.(int)
	if !typeOk {
		logger.Error("could not assert config %v to int", co.name)
		return 0
	}

	return val
}

func (co *configOptions) Int64() int64 {
	val, typeOk := co.rawVal.(int64)
	if !typeOk {
		logger.Error("could not assert config %v to int64", co.name)
		return 0
	}

	return val
}

func (co *configOptions) Float() float32 {
	val, typeOk := co.rawVal.(float32)
	if !typeOk {
		logger.Error("could not assert config %v to float", co.name)
		return 0.0
	}

	return val
}

func (co *configOptions) Bool() bool {
	val, typeOk := co.rawVal.(bool)
	if !typeOk {
		logger.Error("could not assert config %v to bool", co.name)
		return false
	}

	return val
}
