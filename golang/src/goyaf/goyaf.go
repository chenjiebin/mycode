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

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	Log("start server success.")
}
