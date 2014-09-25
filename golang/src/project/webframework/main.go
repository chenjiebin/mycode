package main

import (
	"./controllers"
	"goyaf"
)

func main() {
	goyaf.AddController("/index/index/", &controllers.Index{})
	goyaf.AddController("/index/user/", &controllers.User{})

	goyaf.Run()
}
