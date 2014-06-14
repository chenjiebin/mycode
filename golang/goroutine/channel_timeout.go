package main

import "fmt"

func main() {
	c := make(chan int)
	c := make(chan bool)
	
	go func() {
		for {
			select {
				case v := <- c:
				fmt.Println(v)
				case <- time.After(5 * time.s)
			}
		}
	}
}