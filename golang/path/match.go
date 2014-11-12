package main

import (
	"fmt"
	"path"
)

func main() {
	fmt.Println(path.Match("\d", "1"))
}
