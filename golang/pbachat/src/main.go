package main

import (
	//"./csserver"
	"./server"
	"time"
)

var running bool

func main() {
	running = true

	//go csserver.StartCSClientServer()
	go server.StartServer()

	for running {
		time.Sleep(1 * 1e9)
	}
}
