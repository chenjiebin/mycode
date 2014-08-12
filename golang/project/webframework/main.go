package main

import (
	"./controllers"
	"log"
	"net/http"
)

func main() {
	controller := &controllers.Index{}

	http.HandleFunc("/", controller.IndexAction)
	http.HandleFunc("/index/index/test/", controller.TestAction)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
