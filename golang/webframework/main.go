package main

import (
	"./application/controllers"
	"./conf"
	"flag"
	"git.oschina.net/iceup/goyaf"
)

func main() {
	t := flag.String("type", "master", "start type")
	workerPort := flag.String("workerport", "10001", "worker port")
	env := flag.String("env", "test", "environment")
	flag.Parse()
	goyaf.SetConfig(conf.GetConfigByEnv(*env))

	goyaf.AddController("/index/index/", controllers.Index{})

	if *t == "master" {
		goyaf.StartMaster()
	} else {
		goyaf.StartWorkerListen(*workerPort)
	}
}

func init() {
	goyaf.Log("init main")
}
