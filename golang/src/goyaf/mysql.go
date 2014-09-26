package goyaf

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type GoyafDB struct {
	db *sql.DB
}

var GoyafDBObject *GoyafDB

func (this *GoyafDB) Find(table string, id string) map[string]string {
	rows, err := this.db.Query("SELECT * FROM " + table + " where id=" + id)
	CheckError(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	row := make(map[string]string)
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				row[columns[i]] = string(col.([]byte))
			}
		}
		break
	}
	return row
}

func init() {
	db, err := sql.Open("mysql", "root:@/test?charset=utf8")
	CheckError(err)

	GoyafDBObject = &GoyafDB{db: db}
}
