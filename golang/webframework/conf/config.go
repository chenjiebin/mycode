package conf

var Config map[string]map[string]string

func init() {
	Config = make(map[string]map[string]string)

	//通用配置
	common := make(map[string]string)
	common["mysql-host"] = "127.0.0.1"
	common["mysql-port"] = "3306"
	common["mysql-database"] = "test"
	common["mysql-username"] = "root"
	common["mysql-password"] = ""
	common["mysql-charset"] = "utf8"
	common["mysql-driver_options-1002"] = "SET NAMES utf8mb4"

	common["http-listen-port"] = "10000"

	common["debugmode"] = "0"

	Config["common"] = common

}

func GetConfigByEnv(env string) map[string]string {
	envs, ok := Config[env]
	if !ok {
		return Config["common"]
	}
	return envs
}
