package main

import (
	"log"
)

type Error struct {
	Errno  string
	Errmsg string
}

type Panic struct {
}

func (this *Panic) panicHandle() {
	if r := recover(); r != nil {
		log.Printf("Runtime error caught: %v", r)
	}
}

//func panicHandle() {
//	if r := recover(); r != nil {
//		log.Printf("Runtime error caught: %v", r)
//	}
//}

func main() {
	p := Panic{}

	defer p.panicHandle()
	hello()
}

func hello() {
	panic(Error{Errno: "10020", Errmsg: "cuowu"})
}
