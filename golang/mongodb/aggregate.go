// 演示aggregate

package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
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

	db := session.DB("meiqu618_20150211") //数据库名称
	collection := db.C("shoporders")      //如果该集合已经存在的话，则直接返回

	pipe := collection.Pipe([]bson.M{{"$match": bson.M{"uid": "2016022511161867652"}}})
	iter := pipe.Iter()

	var result []struct {
		Uid string `bson:"uid"`
	}

	fmt.Println(iter.All(&result))
	fmt.Println(result)
}
