package config

import (
	"embed"

	"github.com/tiptophelmet/nomess-core/v2/logger"
	"github.com/tiptophelmet/nomess-core/v2/util"

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

func initFallbackConfigs(fallbackFile embed.FS) *fallbackConfigList {
	if fallbackConfList != nil {
		return fallbackConfList
	}

	var tree *toml.Tree

	if tomlData, err := fallbackFile.ReadFile("config.toml"); err != nil {
		logger.Fatal(err.Error())
	} else if tree, err = toml.Load(string(tomlData)); err != nil {
		logger.Fatal(err.Error())
	} else {
		illegal := util.GetNonIntersecting(getSupportedConfigKeys(), keysRecurse(tree, ""))

		if len(illegal) > 0 {
			logger.Fatal("fallback config.toml has illegal keys: %v", illegal)
		}

	}

	fallbackConfList = &fallbackConfigList{tree.ToMap()}
	return fallbackConfList
}

func keysRecurse(tree *toml.Tree, prefix string) []string {
	var keys []string

	for _, key := range tree.Keys() {
		fullKey := key
		if prefix != "" {
			fullKey = prefix + "." + key
		}

		subTree := tree.Get(key)
		if nestedTree, ok := subTree.(*toml.Tree); ok {
			keys = append(keys, keysRecurse(nestedTree, fullKey)...)
		} else {
			keys = append(keys, fullKey)
		}
	}

	return keys
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

func Init(fallbackFile embed.FS) {
	initFallbackConfigs(fallbackFile)
	initAppConfigs()
}
