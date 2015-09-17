package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		fmt.Println(rand.Intn(5))
	}

	for i := 0; i < 10; i++ {
		//Intn函数返回的最小值为0
		fmt.Println(rand.Intn(1))
	}
}
