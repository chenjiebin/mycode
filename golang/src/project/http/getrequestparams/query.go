package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		queryForm, err := url.ParseQuery(r.URL.RawQuery)
		if err == nil && len(queryForm["id"]) > 0 {
			fmt.Fprintln(w, queryForm["id"][0])
		}
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
