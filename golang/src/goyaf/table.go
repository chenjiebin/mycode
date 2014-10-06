package goyaf

import (
	"database/sql"
	"errors"
	"strings"
)

type Table struct {
	Table   string
	adapter *sql.DB
	where   map[string]string
}

func (this *Table) Select(where map[string]string) map[int]map[string]string {
	if this.adapter == nil {
		this.adapter = getConnect()
	}

	this.Where(where)

	//return make(map[string]string)

	sql := "SELECT * FROM " + this.Table + this.whereToString()
	Debug(sql)

	rows, err := this.adapter.Query(sql)
	CheckError(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	i := 0
	result := make(map[int]map[string]string)
	for rows.Next() {
		row := make(map[string]string)

		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				row[columns[i]] = string(col.([]byte))
			}
		}
		result[i] = row
		i += 1
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
		whereString += k + "=" + v
	}
	return whereString
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

	sql := "INSERT " + this.Table + " SET "
	for k, _ := range data {
		sql += k + "=?,"
	}
	sql = strings.TrimRight(sql, ",")
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

	LastInsertId, err = res.LastInsertId()

	return LastInsertId, err
}
