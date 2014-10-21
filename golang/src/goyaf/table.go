package goyaf

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

//表
type Table struct {
	Table   string
	adapter *sql.DB
	where   map[string]string
}

//查询参数
type Select struct {
	Columns []string
	Where   map[string]string
	Order   map[string]string
	Count   int
	Offset  int
	Group   []string
}

//查询数据
func (this *Table) Select(slt Select) []map[string]string {
	if this.adapter == nil {
		this.adapter = getConnect()
	}

	this.Where(slt.Where)

	//return make(map[string]string)
	sql := fmt.Sprintf("SELECT * FROM %s %s LIMIT %d,%d",
		this.Table,
		this.whereToString(),
		slt.Offset,
		slt.Count)
	Debug(sql)

	rows, err := this.adapter.Query(sql)
	CheckError(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	var result []map[string]string
	for rows.Next() {
		row := make(map[string]string)

		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				row[columns[i]] = string(col.([]byte))
			}
		}

		result = append(result, row)
	}
	return result
}

func (this *Table) Where(where map[string]string) {
	this.where = where
}

func (this *Table) whereToString() string {
	if len(this.where) == 0 {
		return ""
	}

	whereString := " where "
	for k, v := range this.where {
		if strings.IndexAny(k, "=><") == -1 {
			whereString += k + "=" + v + " and "
			continue
		}
		whereString += strings.Replace(k, "?", v, -1) + " and "
	}
	return strings.TrimRight(whereString, " and ")
}

func (this *Table) Order() {

}

func (this *Table) Group() {

}

func (this *Table) Insert(data map[string]string) (LastInsertId int64, err error) {
	if len(data) == 0 {
		return 0, errors.New("data is empty")
	}

	if this.adapter == nil {
		this.adapter = getConnect()
	}

	//todo:这里要采用结构化的方式
	sql := "INSERT " + this.Table + " SET "
	for k, v := range data {
		sql += k + "='" + v + "',"
	}
	sql = strings.TrimRight(sql, ",")
	Debug(sql)

	res, err := this.adapter.Exec(sql)

	//values := make([]interface{}, len(data))
	//i := 0
	//for _, v := range data {
	//	values[i] = v
	//	i += 1
	//}
	//Debug(values)

	//this.GetInsertSql(data)

	//res, err := stmt.Exec(values...)

	LastInsertId, err = res.LastInsertId()

	return LastInsertId, err
}

//func (this *Table) GetInsertSql(data map[string]string) {
//	sql := "INSERT " + this.Table + " SET "
//	for k, v := range data {
//		sql += k + "='" + v + "',"
//	}
//	sql = strings.TrimRight(sql, ",")
//	Debug(sql)
//}

//更新数据
func (this *Table) Update(data map[string]string, where map[string]string) (affect int64, err error) {
	if len(data) == 0 {
		return 0, errors.New("data is empty")
	}

	if this.adapter == nil {
		this.adapter = getConnect()
	}

	sql := "UPDATE " + this.Table + " SET "
	for k, _ := range data {
		sql += k + "=?,"
	}
	sql = strings.TrimRight(sql, ",")

	this.Where(where)

	sql += this.whereToString()
	Debug(sql)

	stmt, err := this.adapter.Prepare(sql)

	values := make([]interface{}, len(data))
	i := 0
	for _, v := range data {
		values[i] = v
		i += 1
	}
	Debug(values)

	res, err := stmt.Exec(values...)

	affect, err = res.RowsAffected()

	return affect, err
}
