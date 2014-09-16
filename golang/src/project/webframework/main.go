package main

import (
	"./controllers"
	"goyaf"
)

func main() {

	controller := &controllers.Index{}
	goyaf.AddController("/index/index/", controller)
	goyaf.AddRoute("/index/index/test/", controller.TestAction)

	goyaf.Run()
}
