package mysql

import (
	"goyaf"
)

type Abstract struct {
}

func (this *Abstract) Find(table string, id string) map[string]string {
	return goyaf.GoyafDBObject.Find(table, id)
}
