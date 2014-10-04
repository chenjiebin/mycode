package mysql

import (
	"goyaf"
)

type Base struct {
}

func (this *Base) Find(table string, id string) map[string]string {
	return goyaf.GoyafDBObject.Find(table, id)
}
