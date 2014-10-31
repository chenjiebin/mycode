package mysql

import (
	"git.oschina.net/iceup/goyaf"
)

type Base struct {
	table      string
	primaryKey string
}

type Mysql interface {
	Find(id string) map[string]string
	Insert(data map[string]string) int64
	Update(data map[string]string, where map[string]string) int64
	FetchAll(slt goyaf.Select) []map[string]string
}

//查询单条记录
func (this *Base) Find(id string) map[string]string {
	tableGateway := &goyaf.Table{Table: this.table}

	result := tableGateway.Select(goyaf.Select{Where: map[string]string{this.primaryKey: id}})
	if len(result) == 0 {
		return make(map[string]string)
	}
	return result[0]
}

//插入数据
func (this *Base) Insert(data map[string]string) int64 {
	tableGateway := &goyaf.Table{Table: this.table}
	result, _ := tableGateway.Insert(data)
	return result
}

//查询单条记录
func (this *Base) Update(data map[string]string, where map[string]string) int64 {
	tableGateway := &goyaf.Table{Table: this.table}

	result, _ := tableGateway.Update(data, where)
	return result
}

//查询列表数据
func (this *Base) FetchAll(slt goyaf.Select) []map[string]string {
	tableGateway := &goyaf.Table{Table: this.table}
	result := tableGateway.Select(slt)

	return result
}
