// 看看mongodb中的timestamp
package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

func main() {
	now := time.Now().Unix()
	fmt.Println(now)
	// output: 1459301711

	now32 := now << 32
	fmt.Println(now32)
	// output: 6267653123741843456

	mongodbnow := bson.MongoTimestamp(now32)
	fmt.Println(mongodbnow)
	// output: 6267653123741843456
}
