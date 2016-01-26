package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("fmt println")
	log.Println("log println")
	fmt.Fprintln(os.Stderr, "fmt fprintln stderr")
	fmt.Fprintln(os.Stdout, "fmt fprintln stdout")
}
