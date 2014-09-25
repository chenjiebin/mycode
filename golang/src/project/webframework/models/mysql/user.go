package mysql

import (
	"fmt"
)

type User struct {
}

func (this *User) Find(id string) {
	fmt.Println(id)
}
