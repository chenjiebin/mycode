package goyaf

var Config map[string]string

func SetConfig(config map[string]string) {
	Config = config

	//Debug(Config)
}

func GetConfigByKey(key string) string {
	return Config[key]
}
