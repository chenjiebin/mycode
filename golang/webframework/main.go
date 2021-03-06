package main

import (
	"./application/controllers"
	"./conf"
	"git.oschina.net/iceup/goyaf"
)

func main() {
	goyaf.SetConfig(conf.Config)
	goyaf.AddController("/index/index/", controllers.Index{})

	goyaf.Run()
}

func init() {
	goyaf.Log("init main")
}
