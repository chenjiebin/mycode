package main

import (
	"os"
	"path/filepath"
)

func main() {
	filepath.Walk("/home/leo",
		func(path string, f os.FileInfo, err error) error {
			if f == nil {
				return err
			}
			if f.IsDir() {
				return nil
			}
			println(path)
			return nil
		})
}
