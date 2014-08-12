package main

import (
	"fmt"
)

const (
	longjourney = 1000
)

type Life struct {
	journey int
	thing   *Thing
}

type Thing struct {
	result string
}

func (life *Life) think(thing *Thing) bool {
	life.thing = thing
	fmt.Println("just think")
	return true
}

func (life *Life) doit() string {
	fmt.Println("just do it")
	return "fail"
	//life.thing.result =
}

func (life *Life) wentByWentBy() {
	life.journey--
}

func (life *Life) end() bool {
	return life.journey == 0
}

func main() {
	life := &Life{journey: longjourney}
	for {
		life.think(&Thing{})
		life.doit()
		if life.thing.result == "fail" {
			life.Go
		}

		life.wentByWentBy()

		if life.end() {
			break
		}

	}

}
