package scheduler

import (
	"fmt"
	"net/http"
	"strings"
)

type Scheduler struct {
	UrlChan chan string
	ReqChan chan *http.Request //请求对象通道
}

func (this *Scheduler) Run() {
	fmt.Println("scheduler run")
	for {
		url := <-this.UrlChan
		req, err := http.NewRequest("GET", url, strings.NewReader(""))
		if err != nil {
			fmt.Println("err: ", err)
			continue
		}
		this.ReqChan <- req
	}

}
