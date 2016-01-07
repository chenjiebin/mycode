package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func main() {
	urls := [100]string{}
	for i := 0; i < 100; i++ {
		urls[i] = "http://so.tv.sohu.com/list_p1122_p2122106_p3_p40_p5_p6_p73_p8_p9_p10" + strconv.Itoa(i+1) + "_p11_p12_p13.html"
	}

	//dowload response
	respChan := make(chan *http.Response, 100)
	//proc result
	procRsChan := make(chan string, 100)

	var down = &Downloader{respChan: respChan}
	for i := 0; i < len(urls); i++ {
		go down.Download(urls[i])
	}

	var proc = &Proc{respChan: respChan, procRsChan: procRsChan}
	go proc.Run()

	var pipeline = &Pipeline{procRsChan: procRsChan}
	go pipeline.Run()

	for {
		time.Sleep(1 * 1e9)
	}

}
