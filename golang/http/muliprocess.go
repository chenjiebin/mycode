package main

import (
	"fmt"
	"log"
	"net/http"
	"net/rpc"
	"os/exec"
	"time"
)

var pid int

type Args struct {
	W http.ResponseWriter
	R *http.Request
}

//默认路由
type Mux struct{}

func (p *Mux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world 2")

	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")

	if err != nil {
		log.Fatal("dialing:", err)
	}
	args := Args{w, r}
	var reply int
	err = client.Call("Arith.ServeHTTP", args, &reply)
	fmt.Println(err)
	fmt.Println(reply)

}

func main() {
	go startByServer()
	go startManage()
	for {
		time.Sleep(1 * 1e9)
	}
}

func startManage() {
	c := exec.Command("/Users/chenjiebin/Sites/mycode/golang/http/simple_http_server")
	c.Run()
}

func startByServer() {
	mux := &Mux{}

	err := http.ListenAndServe(":9091", mux)
	if err != nil {
		fmt.Println(err)
	}
}
