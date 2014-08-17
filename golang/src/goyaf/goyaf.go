package goyaf

import (
	"log"
	"net/http"
)

//goyaf instance
type goyaf struct {
	version string
}

type MyMux struct{}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

type Controller struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func Run() {
	mux := &MyMux{}

	err := http.ListenAndServe(":9090", mux)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
