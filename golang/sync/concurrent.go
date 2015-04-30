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

//测试命令：siege -c 1000 -b -r 10 'http://192.168.3.233:9090/calc'

//var i int32 = 5

var i int = 5000

func calc(w http.ResponseWriter, r *http.Request) {
	//
	if i <= 0 {
		return
	}
	time.Sleep(3000000)

	i = i - 1

	//4200, 4600, 4500,4694.84,4761.90,4694.84,4672.90
	//atomic.AddInt32(&i, 1)
}

func geti(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(i)))

	//w.Write([]byte(strconv.Itoa(int(i))))
}

func main() {
	http.HandleFunc("/calc", calc)
	http.HandleFunc("/geti", geti)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
