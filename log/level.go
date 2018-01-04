package log

import "thelark.cn/golang.ext/color/console"

type __LogLevel__ uint // log 级别

const (
	OFF   __LogLevel__ = iota + 1 // 1
	FATAL                         // 2
	ERROR                         // 3
	WARN                          // 4
	INFO                          // 5
	DEBUG                         // 6
	ALL                           // 7
)

var __LogLevelName__ = map[__LogLevel__]string{
	OFF:   "OFF",
	FATAL: "FATAL",
	ERROR: "ERROR",
	WARN:  "WARN",
	INFO:  "INFO",
	DEBUG: "DEBUG",
	ALL:   "ALL",
}

var __LogLevelColorPrintfFunc__ = map[__LogLevel__]func(format string, args ...interface{}){
	FATAL: console.Magenta().Printf,
	ERROR: console.Red().Printf,
	WARN:  console.Yellow().Printf,
	INFO:  console.Cyan().Printf,
	DEBUG: console.Green().Printf,
	ALL:   console.White().Printf,
}
