package wrong

import (
	"log"
	"thelark.cn/golang.ext/color/console"
	"runtime"
	"fmt"
	"thelark.cn/golang.ext/path"
)

func init() {
	log.SetFlags(0)
}
func core() (string, error) {
	_, file, line, _ := runtime.Caller(2)
	refPath, err := path.GetRelGoPathFromAbsPath(file)
	return fmt.Sprintf("%s: %d", refPath, line), err
}

const __PREFIX__ = "[ERROR] "

var __RED__ = console.Red()

func __printf__(format string, args ...interface{}) {
	__RED__.Print(__PREFIX__)
	log.Printf(format, args...)
}
func __panicf__(format string, args ...interface{}) {
	__RED__.Print(__PREFIX__)
	log.Panicf(format, args...)
}
func __fatalf__(format string, args ...interface{}) {
	__RED__.Print(__PREFIX__)
	log.Fatalf(format, args...)
}
