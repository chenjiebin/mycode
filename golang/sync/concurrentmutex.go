package main

import (
	//"fmt"
	"log"
	"net/http"
	//"sync/atomic"
	//"strings"
	"strconv"
	"sync"
	"time"
)

var mutex sync.Mutex

var i int = 5000

func calc(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()
	if i <= 0 {
		return
	}
	time.Sleep(3000000)
	i = i - 1
}

func geti(w http.ResponseWriter, r *http.Request) {
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
