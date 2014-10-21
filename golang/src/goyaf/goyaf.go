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

//panic处理函数
var panicHandle func()

func SetPanicHandle(f func()) {
	panicHandle = f
}
