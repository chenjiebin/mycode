package main

import (
	"flag"
	"fmt"
)

func main() {
	var pwd string

	port := flag.String("port", ":8080", "http listen port")
	user := flag.String("user", "root", "mysql connect user")
	flag.StringVar(&pwd, "password", "123", "mysql connect password")

	flag.Parse()

	fmt.Println(*port)
	fmt.Println(*user)
	fmt.Println(pwd)
}
