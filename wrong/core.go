package wrong

import (
	"thelark.cn/golang.ext/color/cmd"
	"log"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

const __PREFIX__ = "[ERROR] "

var __RED__ = cmd.Red()

func __printf__(format string, args ...interface{}) {
	__RED__.Printf(__PREFIX__)
	log.Printf(format, args...)
}
