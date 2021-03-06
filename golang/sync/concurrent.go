package main

import (
	//"fmt"
	"log"
	"net/http"
	//"sync/atomic"
	//"strings"
	"strconv"
	"time"
)

var i int = 5000

func calc(w http.ResponseWriter, r *http.Request) {
	if i <= 0 {
		w.Write([]byte("0"))
		return
	}
	//中间的一些校验操作
	time.Sleep(3000000)

	//if i <= 0 {
	//	w.Write([]byte("0"))
	//	return
	//}
	i = i - 1

	//这里可以处理一些后续的功能
	time.Sleep(3000000)
	w.Write([]byte("1"))
}

func geti(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(i)))
}

func main() {
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/geti", geti)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
