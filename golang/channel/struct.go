package main

import (
	"fmt"
	"time"
)

type Player struct {
	Reader chan int
}

func (this *Player) Read() {
	for {
		select {
		case data := <-this.Reader:
			fmt.Println(data)
		}
	}
}

func main() {
	player := &Player{
		Reader: make(chan int, 16),
	}
	go player.Read()

	l(player.Reader)

	for {
		time.Sleep(1 * 1e9)
	}

}

func l(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
}
