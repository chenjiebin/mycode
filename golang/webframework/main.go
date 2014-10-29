package main

import (
	"./application/controllers"
	"./conf"
	"flag"
	"fmt"
	"git.oschina.net/iceup/goyaf"
)

func main() {
	env := flag.String("env", "test", "environment")
	flag.Parse()
	goyaf.SetConfig(conf.GetConfigByEnv(*env))

	goyaf.AddController("/index/index/", controllers.Index{})

	goyaf.Run()
}

func init() {
	fmt.Println("init main")
}
