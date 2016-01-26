package spider

import (
	"../downloader"
	"../pipeline"
	"../processer"
	"../scheduler"
	"net/http"
	"time"
)

func NewSpider() *Spider {
	s := &Spider{}
	s.UrlChan = make(chan string, 100)
	return s
}

type Spider struct {
	UrlChan chan string
}

func (this *Spider) Run() {
	reqChan := make(chan *http.Request, 100)
	respChan := make(chan *http.Response, 100)

	var s = &scheduler.Scheduler{UrlChan: this.UrlChan, ReqChan: reqChan}
	go s.Run()

	var d = &downloader.Downloader{ReqChan: reqChan, RespChan: respChan}
	go d.Run()

	var p = &processer.Processer{RespChan: respChan}
	go p.Run()

	var pl = &pipeline.Pipeline{}
	go pl.Run()

	for {
		time.Sleep(1 * 1e9)
	}
}

func (this *Spider) AddUrl(url string) {
	this.UrlChan <- url
}
