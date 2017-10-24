package wrong

import (
	"thelark.cn/golang.ext/color/console"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

const __PREFIX__ = "[ERROR] "

var __RED__ = console.Red()

func __printf__(format string, args ...interface{}) {
	__RED__.Printf(__PREFIX__)
	log.Printf(format, args...)
}
func __panicf__(format string, args ...interface{}) {
	__RED__.Printf(__PREFIX__)
	log.Panicf(format, args...)
}
func __fatalf__(format string, args ...interface{}) {
	__RED__.Printf(__PREFIX__)
	log.Fatalf(format, args...)
}
