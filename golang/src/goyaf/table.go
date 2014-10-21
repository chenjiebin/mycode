package goyaf

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

//表结构
type Table struct {
	Table   string //表名
	adapter *sql.DB
}

//查询参数
type Select struct {
	Columns string
	Where   map[string]string
	Order   string
	Count   int
	Offset  int
	Group   string
}

func (this *Select) columnsToString() string {
	if len(this.Columns) == 0 {
		return "*"
	}
	return this.Columns
}

func (this *Select) orderToString() string {
	if len(this.Order) == 0 {
		return ""
	}
	return "order by " + this.Order
}

func (this *Select) groupToString() string {
	if len(this.Group) == 0 {
		return ""
	}
	return "group by " + this.Group
}

func (this *Select) countToString() string {
	if this.Count == 0 {
		return ""
	}
	return fmt.Sprintf("limit %d,%d", this.Offset, this.Count)
}

func (this *Select) whereTostring() string {
	if len(this.Where) == 0 {
		return ""
	}

	whereString := " where "
	for k, v := range this.Where {
		if strings.IndexAny(k, "=><?") == -1 {
			whereString += k + "=" + v + " and "
			continue
		}
		whereString += strings.Replace(k, "?", v, -1) + " and "
	}
	return strings.TrimRight(whereString, " and ")
}

//查询数据
func (this *Table) Select(slt Select) []map[string]string {
	if this.adapter == nil {
		this.adapter = getConnect()
	}
	sql := this.SelectToSql(slt)

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

func (this *Table) SelectToSql(slt Select) string {
	sql := fmt.Sprintf("SELECT %s FROM %s %s %s %s %s",
		slt.columnsToString(),
		this.Table,
		slt.whereTostring(),
		slt.groupToString(),
		slt.orderToString(),
		slt.countToString())
	Debug(sql)
	return sql
}

//将where转换为sql的where语句
func (this *Table) whereToString(where map[string]string) string {
	if len(where) == 0 {
		return ""
	}

	whereString := " where "
	for k, v := range where {
		if strings.IndexAny(k, "=><") == -1 {
			whereString += k + "=" + v + " and "
			continue
		}
		whereString += strings.Replace(k, "?", v, -1) + " and "
	}
	return strings.TrimRight(whereString, " and ")
}

//插入数据
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

	sql += this.whereToString(where)
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
