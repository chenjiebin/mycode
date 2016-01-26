package processer

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Processer struct {
	RespChan chan *http.Response
}

func (this *Processer) Run() {
	fmt.Println("processer run")
	for {
		resp := <-this.RespChan
		this.processer(resp)
	}
}

func (this *Processer) processer(resp *http.Response) {
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}

	fmt.Println(string(bodyBytes))
}
