package main

import (
	"./controllers"
	"goyaf"
)

func main() {
	controller := &controllers.Index{}
	goyaf.AddRoute("/index/index/test/", controller.TestAction)

	goyaf.Run()
}
