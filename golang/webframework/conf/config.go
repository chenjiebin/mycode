package conf

var Config map[string]map[string]string

func init() {
	devel := make(map[string]string)
	devel["mysql-host"] = "127.0.0.1"
	devel["mysql-port"] = "3306"
	devel["mysql-database"] = ""
	devel["mysql-username"] = "192.168.3.233"
	devel["mysql-password"] = "192.168.3.233"
	devel["mysql-charset"] = "192.168.3.233"
	devel["mysql-driver_options-1002"] = "SET NAMES utf8mb4"

	Config["devel"] = devel
}
