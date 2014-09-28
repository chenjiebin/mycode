package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Fprintln(w, r.Form["id"])
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {

	}
}
