package log

import (
	"log"
)

func __flags__() int {
	return log.Flags()
}
func __setFlags__(flags int) {
	log.SetFlags(flags)
}
func __prefix__() string {
	return log.Prefix()
}
func __setPrefix__(prefix string) {
	log.SetPrefix(prefix)
}
func __level__(level __LogLevel__) {
	__LogLevelColorPrintfFunc__[level]("[%s]\t", __LogLevelName__[level])
}

func __print__(v ...interface{}) {
	log.Print(v...)
}
func __printf__(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func __println__(v ...interface{}) {
	log.Println(v...)
}
func __fatal__(v ...interface{}) {
	log.Fatal(v...)
}
func __fatalf__(format string, v ...interface{}) {
	log.Fatalf(format, v...)
}
func __fatalln__(v ...interface{}) {
	log.Fatalln(v...)
}
func __panic__(v ...interface{}) {
	log.Panic(v...)
}
func __panicf__(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
func __panicln__(v ...interface{}) {
	log.Panicln(v...)
}
func __output__(calldepth int, s string) error {
	return log.Output(calldepth, s)
}
