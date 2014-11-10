package main

import (
	"log"
)

func a() {
	b()
	c()

	panic("a内错误")
	return
}

func b() {
	panic("b内错误")
}

func c() (err error) {
	panic("c内错误")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()

	a()
}
