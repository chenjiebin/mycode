package main

import (
	"fmt"
)

func main() {
	sql := `INSERT INTO 'log'
VALUES
	(
		"546efd23beb88c1d6fffc589",
		"ios://forum.topicdetail/?id=546efd23beb88c1d6fffc589",
		'{"user_id": "5386ac5589623f3676551cc2","os_v": "ios9.3.1"}',
		1460443075
	) `
	for i := 1; i <= 8041; i++ {
		sql += `,(
		"546efd23beb88c1d6fffc589",
		"ios://forum.topicdetail/?id=546efd23beb88c1d6fffc589",
		'{"user_id": "5386ac5589623f3676551cc2","os_v": "ios9.3.1"}',
		1460443075
	) `
	}
	fmt.Println(sql)
}
