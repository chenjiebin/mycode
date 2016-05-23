// https://godoc.org/github.com/robfig/cron

package main

import (
	"fmt"
	"github.com/robfig/cron"
	"time"
)

func main() {
	c := cron.New()
	c.AddFunc("* * * * * *", func() { fmt.Println("Every second") })
	c.AddFunc("@every 1s", func() { fmt.Println("Every second 2") })
	c.AddFunc("0 30 * * * *", func() { fmt.Println("Every hour on the half hour") })
	c.AddFunc("@hourly", func() { fmt.Println("Every hour") })
	c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty") })
	c.Start()

	for {
		time.Sleep(1 * 1e9)
	}

}
