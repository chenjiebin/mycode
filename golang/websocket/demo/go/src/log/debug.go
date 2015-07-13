package log

import (
	"log"
	"os"
)

var DebugLog *log.Logger

func Debug(v ...interface{}) {
	DebugLog.Println(v...)
}

func init() {
	DebugLog = log.New(os.Stdout, "[DEBUG]", log.LstdFlags)
}
