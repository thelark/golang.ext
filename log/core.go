package log

import (
	"os"
	"log"
	"thelark.cn/golang.ext/wrong"
	"io"
)

type Logger struct {
	level  __LogLevel__ // 日志级别
	flag   bool         // 是否开启记录日志
	logger *log.Logger
}

// TODO: 初始化日志文件
func __init__(fileName string) *Logger {

	//fi, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	//wrong.Println(err)
	//defer fi.Close()
	fii, err := os.Create(fileName)
	wrong.Println(err)
	//defer fii.Close()

	return &Logger{DEBUG, true, log.New(fii, "", Ldate|Ltime|Lshortfile)}
}

func (l *Logger) __setOutput__(w io.Writer) {
	l.logger.SetOutput(w)
}
func (l *Logger) __output__(calldepth int, s string) error {
	return l.logger.Output(calldepth, s)
}

func (l *Logger) __flags__() int {
	return l.logger.Flags()
}
func (l *Logger) __setFlags__(flags int) {
	l.logger.SetFlags(flags)
}
func (l *Logger) __prefix__() string {
	return l.logger.Prefix()
}
func (l *Logger) __setPrefix__(prefix string) {
	l.logger.SetPrefix(prefix)
}

func (l *Logger) __setFile__(fileName string) {
	fi, err := os.OpenFile(fileName, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	defer fi.Close()
	wrong.Println(err)
	l.__setOutput__(fi)
}
func (l *Logger) __setLevel__(level __LogLevel__) {
	l.level = level
}

// TODO: 开启日志记录(文件操作)
func (l *Logger) __open__() {
	l.flag = true
}

// TODO: 关闭日志记录(文件操作)
func (l *Logger) __close__() {
	l.flag = false
}

// TODO: 验证是否开启日志记录(文件操作)
func (l *Logger) __check__(level __LogLevel__) bool {
	if l.flag {
		return false
	}
	if l.level < level {
		return false
	}
	return true
}
func (l *Logger) __level__(level __LogLevel__) {
	l.logger.Printf("[%s]\t", __LogLevelName__[level])
}

func (l *Logger) __printf__(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}
func (l *Logger) __print__(v ...interface{}) {
	l.logger.Print(v...)
}
func (l *Logger) __println__(v ...interface{}) {
	l.logger.Println(v...)
}
func (l *Logger) __fatal__(v ...interface{}) {
	l.logger.Fatal(v...)
}
func (l *Logger) __fatalf__(format string, v ...interface{}) {
	l.logger.Fatalf(format, v...)
}
func (l *Logger) __fatalln__(v ...interface{}) {
	l.logger.Fatalln(v...)
}
func (l *Logger) __panic__(v ...interface{}) {
	l.logger.Panic(v...)
}
func (l *Logger) __panicf__(format string, v ...interface{}) {
	l.logger.Panicf(format, v...)
}
func (l *Logger) __panicln__(v ...interface{}) {
	l.logger.Panicln(v...)
}
