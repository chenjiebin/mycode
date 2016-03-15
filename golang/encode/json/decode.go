package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsons = []byte(`[
		{"name": "Tom", "age": 18},
		{"name": "Anny", "age": 16}
	]`)
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	var users []User
	err := json.Unmarshal(jsons, &users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(users)
	// output: [{Tom 18} {Anny 16}]
}
