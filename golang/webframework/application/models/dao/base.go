package dao

import (
	"../mysql"
	"git.oschina.net/iceup/goyaf"
)

type Base struct {
	mysql mysql.Mysql
}

//查找单条数据
func (this *Base) Find(id string) map[string]string {
	return this.mysql.Find(id)
}

//插入数据
func (this *Base) Insert(params map[string]string) int64 {
	return this.mysql.Insert(params)
}

//更新数据
func (this *Base) Update(data map[string]string, where map[string]string) int64 {
	return this.mysql.Update(data, where)
}

//查找数据
func (this *Base) FetchAll(slt goyaf.Select) []map[string]string {
	return this.mysql.FetchAll(slt)
}
