package downloader

import (
	"fmt"
	"net/http"
)

type Downloader struct {
	ReqChan  chan *http.Request  //请求对象通道
	RespChan chan *http.Response //结果集通道
}

func (this *Downloader) Run() {
	fmt.Println("downloader run")
	for {
		req := <-this.ReqChan
		go this.download(req)
	}
}

func (this *Downloader) download(req *http.Request) {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	this.RespChan <- resp
}
