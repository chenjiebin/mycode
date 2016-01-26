package main

import (
	"./spider"
)

func main() {
	s := spider.NewSpider()
	s.AddUrl("http://so.tv.sohu.com/list_p1122_p2122106_p3_p40_p5_p6_p73_p8_p9_p101_p11_p12_p13.html")
	s.Run()
}
