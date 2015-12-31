package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	var down = &Downloader{}
	str, _ := down.Download()
	fmt.Println(str)
}

//调度器
type Sched struct {
}

//下载器
type Downloader struct {
}

//下载页面
func (this *Downloader) Download() (body string, err error) {
	var req *http.Request
	req, err = http.NewRequest("GET", "http://www.163.com/", strings.NewReader(""))
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
