package main

import (
	"fmt"
)

type user struct {
	name string
}

func main() {
	fmt.Println(&user{name: "cjb"})
}
