package main

import (
	"fmt"
	//	"math/rand"
	//	"os"
	"time"
)

func main() {
	//	t := time.Now()

	tmp := make([]byte, 10000000)
	for i, _ := range tmp {
		tmp[i] = 0
	}

	tmp2 := make([]byte, 10000000)
	for i, _ := range tmp2 {
		tmp2[i] = 1
	}

	t := time.Now()
	count := 0
	for i := 0; i < 10000000; i++ {
		if tmp[i]|tmp2[i] == 1 {
			//		if tmp[i] == 1 || tmp2[i] == 1 {
			count = count + 1
		}
	}

	fmt.Println(count)
	fmt.Println(time.Now().Sub(t))

	//	file, err := os.OpenFile("./tmp.txt", os.O_CREATE|os.O_WRONLY, 0600)
	//	if err != nil {
	//		panic(err)
	//	}
	//	n, err := file.Write(tmp)
	//	if err != nil {
	//		panic(err)
	//	}
	//	fmt.Println(n)
}
