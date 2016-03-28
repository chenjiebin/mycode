package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	time.Sleep(1 * 1e9)
	fmt.Println(time.Now().Sub(t))
}
