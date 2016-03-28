// 演示mongodb插入操作
package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var (
	host = "172.16.27.134"
	port = 27017
)

func main() {
	session, err := mgo.Dial("172.16.27.134:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	db := session.DB("test")

	collection := db.C("user") //如果该集合已经存在的话，则直接返回

	err = collection.Insert(&record{Index: "Ale", Data: {"goods_id": "1", "user_id": "100"}, Time: 1456814139})
	if err != nil {
		panic(err)
	}
}
