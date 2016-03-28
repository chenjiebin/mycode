// 演示插入多层json格式的数据
// {
//    "index": "goods_detail",
//    "goods_id":"1",
//    "user_id":"100",
//    "time": 1456814139
// }

package main

import (
	//	"fmt"

	"gopkg.in/mgo.v2"
	//	"gopkg.in/mgo.v2/bson"
)

type record struct {
	Index string
	Data  interface{}
	Time  int
}

func main() {
	session, err := mgo.Dial("192.168.1.119:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("meiqu_statistics") //数据库名称
	collection := db.C("record")         //如果该集合已经存在的话，则直接返回

	err = collection.Insert(&record{Index: "Ale", Data: {"goods_id": "1", "user_id": "100"}, Time: 1456814139})
	if err != nil {
		panic(err)
	}
}
