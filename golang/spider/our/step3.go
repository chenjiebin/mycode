//使用多协程方式进行下载
//增加通道将下载后的数据传递出来
package main

import (
	"fmt"
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

	respChan := make(chan *http.Response, 100)

	var down = &Downloader{respChan: respChan}
	for i := 0; i < len(urls); i++ {
		go down.Download(urls[i])
	}

	var proc = &Proc{respChan: respChan}
	go proc.Run()

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
	respChan chan *http.Response //结果集通道
}

//启动页面分析器
func (this *Proc) Run() {
	for {
		//等待通道传递数据，如果没有数据则会处于阻塞状态
		resp := <-this.respChan
		this.Proc(resp)
	}
}

//具体分析方法
func (this *Proc) Proc(resp *http.Response) {
	fmt.Println(resp.Request.URL.String(), "proc")
}

//可能存在的问题
//当正在下载的时候，我们需要中断服务重新启动，希望程序友好的退出。
