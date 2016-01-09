package conf

var Config map[string]map[string]string

func init() {
	Config = make(map[string]map[string]string)

	//通用配置
	common := make(map[string]string)
	common["mysql-host"] = "192.168.1.102"
	common["mysql-port"] = "3306"
	common["mysql-database"] = "test"
	common["mysql-username"] = "root"
	common["mysql-password"] = ""
	common["mysql-charset"] = "utf8"
	common["mysql-driver_options-1002"] = "SET NAMES utf8mb4"

	common["http-listen-port"] = "10000"

	common["debugmode"] = "1"

	Config["common"] = common

	//生产环境配置
	product := make(map[string]string)
	product["stdout"] = "/root/go/log.log"
	product["stderr"] = "/root/go/error.log"
	Config["product"] = product

}
