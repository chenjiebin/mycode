package goyaf

var Config map[string]map[string]string

func SetConfig(config map[string]map[string]string) {
	Config = config

	//Debug(Config)
}

func getConfigByKey(key string) string {
	return Config["devel"][key]
}
