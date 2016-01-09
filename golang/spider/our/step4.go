//使用多协程方式进行下载
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
	//url chan
	urlChan := make(chan string, 100)
	//dowload response
	respChan := make(chan *http.Response, 100)
	//quit chan
	downQuitChan := make(chan int, 1)
	//quit chan
	procQuitChan := make(chan int, 1)

	var down = &Downloader{urlChan: urlChan, respChan: respChan, quitChan: downQuitChan}
	go down.Run()

	var proc = &Proc{respChan: respChan, quitChan: procQuitChan}
	go proc.Run()

	urls := [100]string{}
	for i := 0; i < 100; i++ {
		urls[i] = "http://so.tv.sohu.com/list_p1122_p2122106_p3_p40_p5_p6_p73_p8_p9_p10" + strconv.Itoa(i+1) + "_p11_p12_p13.html"
	}
	for _, url := range urls {
		urlChan <- url
	}

	//休眠10秒等待下载和处理完成
	//	time.Sleep(10 * 1e9)

	//退出下载协程和页面分析协程
	downQuitChan <- 1
	procQuitChan <- 1

	//看看退出后urlChan和respChan数量
	for i := 0; i < 5; i++ {
		fmt.Println("sleeping...")
		time.Sleep(1 * 1e9)
	}
	fmt.Println("urlChan len", len(urlChan))
	fmt.Println("respChan len", len(respChan))

	for {
		time.Sleep(1 * 1e9)
	}
}

//下载器
type Downloader struct {
	urlChan  chan string         //url通道
	respChan chan *http.Response //结果集通道
	quitChan chan int            //接收退出消息
}

//启动下载器
func (this *Downloader) Run() {
OuterLoop:
	for {
		select {
		case url := <-this.urlChan:
			go this.Download(url)
		case <-this.quitChan:
			break OuterLoop
		}
	}

	fmt.Println("downloader exit")
}

//下载页面
func (this *Downloader) Download(url string) (err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", url, strings.NewReader(""))
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	client := &http.Client{}
	var resp *http.Response
	resp, err = client.Do(req)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(resp.Request.URL.String(), "download")
	this.respChan <- resp
	return
}

//页面分析器
type Proc struct {
	respChan chan *http.Response //结果集通道
	quitChan chan int            //接收退出消息
}

//启动页面分析器
func (this *Proc) Run() {
OuterLoop:
	for {
		select {
		case resp := <-this.respChan:
			this.proc(resp)
		case <-this.quitChan:
			break OuterLoop
		}
	}
	fmt.Println("proc exit")
}

//具体分析方法
func (this *Proc) proc(resp *http.Response) (err error) {
	fmt.Println(resp.Request.URL.String(), "proc")

	defer resp.Body.Close()

	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	body := string(bodyBytes)
	fmt.Println(resp.Request.URL, len(body))
	return
}
