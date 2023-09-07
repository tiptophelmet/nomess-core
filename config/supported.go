package config

var supportedConfigs map[string]string

func Register(configs map[string]string) {
	supportedConfigs = configs
}

func getSupportedConfigKeys() []string {
	keys := make([]string, 0, len(supportedConfigs))

	for k := range supportedConfigs {
		keys = append(keys, k)
	}

	return keys
}
