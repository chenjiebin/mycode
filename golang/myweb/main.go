package main

import (
	"./application/controllers"
	"./conf"
	"flag"
	"git.oschina.net/iceup/goyaf"
)

func main() {
	env := flag.String("env", "test", "environment")
	flag.Parse()
	goyaf.SetConfig(conf.GetConfigByEnv(*env))

	goyaf.AddController("/index/index/", controllers.Index{})
	goyaf.AddController("/index/user/", controllers.User{})

	goyaf.Run()
}
