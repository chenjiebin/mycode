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

	var down = &Downloader{urlChan: urlChan, respChan: respChan}
	go down.Run()

	var proc = &Proc{respChan: respChan}
	go proc.Run()

	urls := [100]string{}
	for i := 0; i < 100; i++ {
		urls[i] = "http://so.tv.sohu.com/list_p1122_p2122106_p3_p40_p5_p6_p73_p8_p9_p10" + strconv.Itoa(i+1) + "_p11_p12_p13.html"
	}
	for _, url := range urls {
		urlChan <- url
	}

	for {
		time.Sleep(1 * 1e9)
	}

}

type Sche struct {
}

//下载器
type Downloader struct {
	urlChan  chan string         //url通道
	respChan chan *http.Response //结果集通道
}

//启动页面分析器
func (this *Downloader) Run() {
	for {
		url := <-this.urlChan
		go this.Download(url)
	}
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
	respChan chan *http.Response
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
	fmt.Println(resp.Request.URL, len(body))
	return
}
