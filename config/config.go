package config

import (
	"os"

	"github.com/tiptophelmet/nomess-core/logger"
	"github.com/tiptophelmet/nomess-core/util"

	"github.com/pelletier/go-toml"
)

type fallbackConfigList struct {
	list map[string]interface{}
}

func fallback(configName string) interface{} {
	if val, present := fallbackConfList.list[configName]; present {
		return val
	} else {
		return nil
	}
}

var fallbackConfList *fallbackConfigList

func initFallbackConfigs() *fallbackConfigList {
	if fallbackConfList != nil {
		return fallbackConfList
	}

	var list *toml.Tree

	if tomlData, err := os.ReadFile("../../config.toml"); err != nil {
		logger.Fatal(err.Error())
	} else if list, err = toml.Load(string(tomlData)); err != nil {
		logger.Fatal(err.Error())
	} else {
		illegal := util.GetNonIntersecting(getSupportedConfigKeys(), list.Keys())

		if len(illegal) > 0 {
			logger.Fatal("fallback config.toml has illegal keys: %v", illegal)
		}

	}

	fallbackConfList = &fallbackConfigList{list.ToMap()}
	return fallbackConfList
}

type configList struct {
	list map[string]*env
}

var confList *configList

func initAppConfigs() *configList {
	if confList != nil {
		return confList
	}

	list := make(map[string]*env)

	for configName, envName := range supportedConfigs {
		list[configName] = initEnv(envName, fallback(configName))
	}

	confList = &configList{list}
	return confList
}

func Init() {
	initFallbackConfigs()
	initAppConfigs()
}
