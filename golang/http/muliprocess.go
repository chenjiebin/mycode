package main

import (
	"fmt"
	//"log"
	"net/http"
	"net/http/httputil"
	"net/url"
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
	remote, err := url.Parse("http://127.0.0.1:9090")
	if err != nil {
		panic(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote)
	proxy.ServeHTTP(w, r)

	fmt.Fprintln(w, "Hello world 2")

	//fmt.Println(r)
	//http.
	//resp, err := http.DefaultClient.Do(r)
	//defer resp.Body.Close()
	//if err != nil {
	//	panic(err)
	//}
	//for k, v := range resp.Header {
	//	for _, vv := range v {
	//		w.Header().Add(k, vv)
	//	}
	//}
	//for _, c := range resp.SetCookie {
	//	w.Header().Add("Set-Cookie", c.Raw)
	//}
	//w.WriteHeader(resp.StatusCode)
	//result, err := ioutil.ReadAll(resp.Body)
	//if err != nil && err != os.EOF {
	//	panic(err)
	//}
	//w.Write(result)

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
