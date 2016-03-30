// 从oplog读取数据

package main

import (
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	read()
}

// 从oplog读取数据
func read() {
	session, err := mgo.Dial("172.16.27.134:10001") //连接数据库
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)

	db := session.DB("local")
	c := db.C("oplog.rs")

	result := []struct {
		Timestamp    bson.MongoTimestamp `bson:"ts"`
		HistoryID    int64               `bson:"h"`
		MongoVersion int                 `bson:"v"`
		Operation    string              `bson:"op"`
		Namespace    string              `bson:"ns"`
		Object       bson.M              `bson:"o"`
		QueryObject  bson.M              `bson:"o2"`
	}{}
	c.Find(bson.M{}).All(&result)

	for _, v := range result {
		fmt.Println(v)
	}

}
