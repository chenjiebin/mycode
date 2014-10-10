package main

import (
	"fmt"
)

func main() {
	var users []map[string]string

	users = append(users, map[string]string{"id": "1"}, map[string]string{"id": "2"})

	fmt.Println(users)
}
