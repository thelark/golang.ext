package log

import (
	"log"
	"os"
	"runtime"
)

var (
	__std__   = log.New(os.Stderr, "[FATAL] ", LstdFlags)
	__Fatal__ = log.New(os.Stderr, "[FATAL] ", LstdFlags)
	__Error__ = log.New(os.Stderr, "[ERROR] ", LstdFlags)
	__Warn__  = log.New(os.Stderr, "[WRAN] ", LstdFlags)
	__Info__  = log.New(os.Stderr, "[INFO] ", LstdFlags)
	__Debug__ = log.New(os.Stderr, "[DEBUG] ", LstdFlags)
	__All__   = log.New(os.Stderr, "[ALL] ", LstdFlags)
)

func Flags() int {
	return __std__.Flags()
}
func SetFlags(flags int) {
	__std__.SetFlags(flags)
	__Fatal__.SetFlags(flags)
	__Error__.SetFlags(flags)
	__Warn__.SetFlags(flags)
	__Info__.SetFlags(flags)
	__Debug__.SetFlags(flags)
}
func Prefix() string {
	return __std__.Prefix()
}
func SetPrefix(prefix string) {
	__std__.SetPrefix(prefix)
	__Fatal__.SetPrefix(prefix)
	__Error__.SetPrefix(prefix)
	__Warn__.SetPrefix(prefix)
	__Info__.SetPrefix(prefix)
	__Debug__.SetPrefix(prefix)
}

func Print(v ...interface{}) {
	log.Print(v...)
}
func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func Println(v ...interface{}) {
	log.Println(v...)
}
func Panic(v ...interface{}) {
	log.Panic(v...)
}
func Panicf(format string, v ...interface{}) {
	log.Panicf(format, v...)
}
func Panicln(v ...interface{}) {
	log.Panicln(v...)
}
func Fatal(v ...interface{}) {
	__level__(FATAL)
	__Fatal__.Fatal(v...)
}
func Fatalf(format string, v ...interface{}) {
	__level__(FATAL)
	__Fatal__.Fatalf(format, v...)
}
func Fatalln(v ...interface{}) {
	__level__(FATAL)
	__Fatal__.Fatalln(v...)
}
func Error(v ...interface{}) {
	_, file, line, ok := runtime.Caller(-1)
	if !ok {
		file = "???"
		line = 0
	}

	//__level__(ERROR)
	log.Print(file, line)
	__Error__.Print(file, line)
	__Error__.Print(v...)
}
func Errorf(format string, v ...interface{}) {
	__level__(ERROR)
	__Error__.Printf(format, v...)
}
func Errorln(v ...interface{}) {
	__level__(ERROR)
	__Error__.Println(v...)
}
func Warn(v ...interface{}) {
	__level__(WARN)
	__Warn__.Print(v...)
}
func Warnf(format string, v ...interface{}) {
	__level__(WARN)
	__Warn__.Printf(format, v...)
}
func Warnln(v ...interface{}) {
	__level__(WARN)
	__Warn__.Println(v...)
}
func Info(v ...interface{}) {
	__level__(INFO)
	__Info__.Print(v...)
}
func Infof(format string, v ...interface{}) {
	__level__(INFO)
	__Info__.Printf(format, v...)
}
func Infoln(v ...interface{}) {
	__level__(INFO)
	__Info__.Println(v...)
}
func Debug(v ...interface{}) {
	__level__(DEBUG)
	__Debug__.Print(v...)
}
func Debugf(format string, v ...interface{}) {
	__level__(DEBUG)
	__Debug__.Printf(format, v...)
}
func Debugln(v ...interface{}) {
	__level__(DEBUG)
	__Debug__.Println(v...)
}

func Output(calldepth int, s string) error {
	return __output__(calldepth, s)
}

func __level__(level __LogLevel__) {
	__LogLevelColorPrintfFunc__[level]("[%s]\t", __LogLevelName__[level])
}

func __output__(calldepth int, s string) error {
	return log.Output(calldepth, s)
}
