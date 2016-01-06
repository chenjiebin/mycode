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

//下载器
type Downloader struct {
	respChan chan *http.Response //结果集通道
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
	respChan   chan *http.Response
	procRsChan chan string
}

//启动页面分析器
func (this *Proc) Run() {
	for {
		resp := <-this.respChan
		this.proc(resp)
	}
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
	this.procRsChan <- body
	return
}

//存储器
type Pipeline struct {
	procRsChan chan string
}

//启动存储器
func (this *Pipeline) Run() {
	for {
		procRs := <-this.procRsChan
		this.pipeline(procRs)
	}
}

//处理页面分析后的结果
func (this *Pipeline) pipeline(procRs string) (err error) {
	fmt.Println(len(procRs))
	return
}

//存在的问题：
//1. 下载页面结果无法返回只能在下载器里处理
//2. 下载协程数量没有控制，我们希望有一些类似线程池的东西，闲置状态开启若干协程，任务繁重的时候开启更多协程，但有一个最大数量
//3. 先有url，然后才开启下载，正常来讲应该是下载器处于监听的状态，监听到有url才进行下载
