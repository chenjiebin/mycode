// golang gc调试
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp := make(map[int]int)
		for i := 0; i < 100000; i++ {
			tmp[i] = i
		}
		fmt.Fprintln(w, "hello", len(tmp))
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
