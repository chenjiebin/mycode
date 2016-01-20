package spider


import (
	"../scheduler"
	"../downloader"
	"../pipeline"
	"../processer"
	"time"
)


type Spider struct {
	
}

func (this *Spider) Run() {
	var s = &scheduler.Scheduler{}
	go s.Run()
	
	var d = &downloader.Downloader{}
	go d.Run()
	
	var p = &processer.Processer{}
	go p.Run()
	
	var pl = &pipeline.Pipeline{}
	go pl.Run()
	
	for {
		time.Sleep(1 * 1e9)
	}
}