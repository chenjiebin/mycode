package main

import (
	"bytes"
	"fmt"
	"strconv"

	elastigo "github.com/mattbaird/elastigo/lib"
)

func main() {
	c := elastigo.NewConn()
	fmt.Println()

	indexer := c.NewBulkIndexer(10)
	// Create a custom Sender Func, to allow inspection of response/error
	indexer.Sender = func(buf *bytes.Buffer) error {
		fmt.Println(buf)
		// @buf is the buffer of docs about to be written
		respJson, err := c.DoCommand("POST", "/_bulk", nil, buf)
		if err != nil {
			// handle it better than this
			fmt.Println(string(respJson))
		}
		return err
	}
	indexer.Start()
	for i := 0; i < 20; i++ {
		err := indexer.Index("twitter", "user", strconv.Itoa(i), "", "", nil, `{"name":"bob"}`, true)
		fmt.Println(err)
	}
	indexer.Stop()
}
