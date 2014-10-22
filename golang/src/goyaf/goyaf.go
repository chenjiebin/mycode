package goyaf

import (
	"log"
	"net/http"
)

type Goyaf struct {
	version string
}

func Run() {
	mux := &GoyafMux{}

	err := http.ListenAndServe(":8002", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	Log("start server success.")
}

var panicHandle interface{}

func SetPanicHandle(c interface{}) {
	panicHandle = c
}
