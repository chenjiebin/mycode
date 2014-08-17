package main

import (
	"./controllers"
	"goyaf"
	"net/http"
)

func main() {
	controller := &controllers.Index{}

	http.HandleFunc("/index/index/index/", controller.IndexAction)
	http.HandleFunc("/index/index/test/", controller.TestAction)

	goyaf.Run()
}
