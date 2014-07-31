package main

import (
	"flag"
	"fmt"
)

var (
	port = flag.String("port", ":8080", "http listen port")
)

func main() {
	flag.Parse()
	fmt.Println(*port)
}
