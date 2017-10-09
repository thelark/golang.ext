package wrong

import (
	"runtime"
	"log"
	"thelark.cn/golang.ext/color/cmd"
	"thelark.cn/golang.ext/path"
)

var red = cmd.Red()

const prefix = "[ERROR] "

func init() {
	log.SetFlags(log.Ldate | log.Ltime)
}
func Println(err error) {
	if err != nil {
		_, filePath, line, _ := runtime.Caller(1)
		red.Printf(prefix)
		log.Printf("%v:%v Error: %v\n", path.GetFileRelPath(filePath), line, err)
	}
}
func Panicln(err error) error {
	if err != nil {
		_, filePath, line, _ := runtime.Caller(1)
		red.Printf("[ERROR] ")
		log.Panicf("%v:%v Error: %v\n", path.GetFileRelPath(filePath), line, err)
		return err
	}

	return nil
}
func Fatalln(err error) error {
	if err != nil {
		_, filePath, line, _ := runtime.Caller(1)
		red.Printf(prefix)
		log.Fatalln("%v:%v Error: %v\n", path.GetFileRelPath(filePath), line, err)
		return err
	}
	return nil
}
