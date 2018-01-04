package log

import (
	"fmt"
	"io"
	"log"
)

type Logger struct {
	level  __LogLevel__ // 日志级别
	flag   bool         // 是否开启记录日志
	logger *log.Logger
}

/**
 * TODO: 初始化日志文件
 * @params writer 文件流
 * @params prefix 日志前缀
 * @params flags 日志 flag
 * @return *Logger
 */
func __init__(out io.Writer, prefix string, flags int) *Logger {
	return &Logger{DEBUG, true, log.New(out, prefix, flags)}
}

/**
 * TODO: 设置输出
 * @params writer 文件流
 */
func (l *Logger) __setOutput__(w io.Writer) {
	l.logger.SetOutput(w)
}
func (l *Logger) __output__(calldepth int, s string) error {
	return l.logger.Output(calldepth, s)
}

/**
 * TODO: 获取 flags
 * @return int
 */
func (l *Logger) __flags__() int {
	return l.logger.Flags()
}

/**
 * TODO: 设置 flags
 * @params flags flags
 */
func (l *Logger) __setFlags__(flags int) {
	l.logger.SetFlags(flags)
}

/**
 * TODO: 获取 前缀
 * @return string
 */
func (l *Logger) __prefix__() string {
	return l.logger.Prefix()
}

/**
 * TODO: 设置 前缀
 * @params prefix 前缀
 */
func (l *Logger) __setPrefix__(prefix string) {
	l.logger.SetPrefix(prefix)
}

/**
 * TODO: 设置 级别
 * @params level 日志级别
 */
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
	return l.flag && l.level >= level
}
func (l *Logger) __level__(level __LogLevel__, v ...interface{}) string {
	return fmt.Sprintf("[%s]\t", __LogLevelName__[level])
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
