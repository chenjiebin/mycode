package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("id")
		if err == nil {
			fmt.Fprintln(w, "Domain:", cookie.Domain)
			fmt.Fprintln(w, "Expires:", cookie.Expires)
			fmt.Fprintln(w, "Name:", cookie.Name)
			fmt.Fprintln(w, "Value:", cookie.Value)
		}

	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
