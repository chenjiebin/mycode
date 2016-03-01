// 当Content-Type为application/json的时候
// 获取body的json格式字符串
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintln(w, k, v)
		}

		r.ParseForm()
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {

		}
		fmt.Fprintln(w, string(body))
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
