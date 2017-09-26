package wrong

import (
	"runtime"
	"log"
	"thelark.cn/golang.ext/color/cmd"
	"thelark.cn/golang.ext/file"
)

var red = cmd.Red()

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}
func Println(err error) {
	if err != nil {
		_, filePath, line, _ := runtime.Caller(1)
		red.Printf("[ERROR] ")
		log.Printf("%v:%v Error: %v\n", file.GetFileRelPath(filePath), line, err)
	}
}
func Panicln(err error) {
	if err != nil {
		_, filePath, line, _ := runtime.Caller(1)
		red.Printf("[ERROR] ")
		log.Panicf("%v:%v Error: %v\n", file.GetFileRelPath(filePath), line, err)
	}
}
func Fatalln(err error) {
	if err != nil {
		_, filePath, line, _ := runtime.Caller(1)
		red.Printf("[ERROR] ")
		log.Fatalln("%v:%v Error: %v\n", file.GetFileRelPath(filePath), line, err)
	}
}
