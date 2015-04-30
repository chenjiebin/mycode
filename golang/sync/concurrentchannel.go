package main

import (
	"log"
	"net/http"
	"strconv"
	"time"
)

var chs chan chan int

var i int = 5000

func calc(w http.ResponseWriter, r *http.Request) {
	if i <= 0 {
		w.Write([]byte("0"))
		return
	}
	//中间的一些校验操作
	time.Sleep(3000000)

	var c chan int
	c = make(chan int, 1)
	chs <- c

	select {
	case r := <-c:
		//如果处理成功
		if r == 1 {
			//处理后续的功能，这里默认用休眠3毫秒来实现
			time.Sleep(3000000)
		}
		w.Write([]byte(strconv.Itoa(r)))
	}
}

func chanSelect(chs chan chan int) {
	for {
		select {
		case c := <-chs:
			if i <= 0 {
				c <- 0
			} else {
				i = i - 1
				c <- 1
			}
		}
	}
}

func geti(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(int(i))))
}

func main() {
	chs = make(chan chan int, 5000)
	go chanSelect(chs)

	http.HandleFunc("/calc", calc)
	http.HandleFunc("/geti", geti)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
