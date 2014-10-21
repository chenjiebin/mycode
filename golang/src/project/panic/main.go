package main

import (
	"log"
)

type Error struct {
	Errno  string
	Errmsg string
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("Runtime error caught: %v", r)
		}
	}()
	hello()
}

func hello() {
	panic(Error{Errno: "10020", Errmsg: "cuowu"})
}
