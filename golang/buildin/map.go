package main

import (
	"fmt"
)

func main() {
	users := make(map[string]string)
	users["1"] = "name"

	fmt.Println(len(users))

	delete(users, "1")
	fmt.Println(len(users))
}
