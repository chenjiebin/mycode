package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	type User struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	users := []User{
		User{Name: "Tom", Age: 18},
		User{Name: "Tom", Age: 18},
	}
	jsons, err := json.Marshal(users)
	fmt.Println(string(jsons), err)
	// output: [{"name":"Tom","age":18},{"name":"Tom","age":18}] <nil>
}
