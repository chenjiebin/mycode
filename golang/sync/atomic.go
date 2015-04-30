package main

import (
	//"fmt"
	"log"
	"net/http"
	"sync/atomic"
	//"strings"
	"strconv"
	//"time"
)

//测试命令：siege -c 1000 -b -r 10 'http://192.168.3.233:9090/calc'

var i int32 = 5

//var i int = 10

func calc(w http.ResponseWriter, r *http.Request) {
	//4255.32, 4600, 4854.37, 4854.37, 4784.69, 4237.29, 4464.29
	//i = i + 1

	//4200, 4600, 4500,4694.84,4761.90,4694.84,4672.90
	atomic.AddInt32(&i, 1)
}

func geti(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte(strconv.Itoa(i)))

	w.Write([]byte(strconv.Itoa(int(i))))
}

func main() {
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/geti", geti)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
