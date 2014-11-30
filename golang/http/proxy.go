package main

import (
	"fmt"
	//"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

//默认路由
type Mux struct {
	workerUrl string
}

func (p *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(p.workerUrl)
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)
}

func main() {
	startByServer()
}

func startByServer() {
	mux := &Mux{}
	mux.workerUrl = "http://127.0.0.1:9090"

	go signalwork(mux)

	err := http.ListenAndServe(":9091", mux)
	if err != nil {
		fmt.Println(err)
	}
}

func signalwork(mux *Mux) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGUSR2)
	select {
	case s := <-ch:
		mux.workerUrl = "http://127.0.0.1:9092"
		fmt.Println("Got signal:", s)
	}
}
