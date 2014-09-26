package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			return
		}

		r.ParseMultipartForm(32 << 20)
		values := r.MultipartForm.Value["id"]
		if len(values) > 0 {
			fmt.Fprintln(w, values[0])
		}

	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
