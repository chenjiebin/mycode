package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
)

func main() {
	startByServer()
}

func normal() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world")
	})

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func startByServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r.Cookies())
		fmt.Fprintf(w, "Hello world")
	})
	http.HandleFunc("/app/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(r)
		fmt.Fprintf(w, "Hello app")
	})

	l, err := net.Listen("tcp", ":15000")
	if nil != err {
		log.Fatalln(err)
	}

	err = http.Serve(l, nil)
	if err != nil {
		log.Fatalln(err)
	}
}
