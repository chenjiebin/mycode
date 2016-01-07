//使用单线程方式进行下载
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	urls := [100]string{}
	for i := 0; i < 100; i++ {
		urls[i] = "http://so.tv.sohu.com/list_p1122_p2122106_p3_p40_p5_p6_p73_p8_p9_p10" + strconv.Itoa(i+1) + "_p11_p12_p13.html"
	}

	fmt.Println("start")

	var down = &Downloader{}
	for i := 0; i < len(urls); i++ {
		down.Download(urls[i])
		fmt.Println(urls[i])
	}

	fmt.Println("finish")
}

//下载器
type Downloader struct {
}

//下载页面
func (this *Downloader) Download(url string) (body string, err error) {
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

	defer resp.Body.Close()

	var bodyBytes []byte
	bodyBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	body = string(bodyBytes)
	return
}

//存在的问题
//单协程下载，下载慢
