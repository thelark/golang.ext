package log

import (
	"log"
	"io"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func New(writer io.Writer, prefix string, flags int) *Logger {
	return __init__(writer, prefix, flags)
}

// 开启记录日志
func (l *Logger) Open() {
	l.__open__()
}

// 关闭记录日志
func (l *Logger) Close() {
	l.__close__()
}

func (l *Logger) SetFlags(flags int) {
	l.__setFlags__(flags)
}
func (l *Logger) SetOutput(writer io.Writer) {
	l.__setOutput__(writer)
}
func (l *Logger) SetLevel(level __LogLevel__) {
	l.__setLevel__(level)
}

//FATAL => Magenta // 致命
//ERROR => Red // 错误
//WARN => Yellow // 警告
//INFO=> Cyan // 信息
//DEBUG => Green // 调试
//ALL => White // 所有
func (l *Logger) Print(v ...interface{}) {
	if !l.__check__(ALL) {
		return
	}
	l.__print__(v...)
}
func (l *Logger) Printf(format string, v ...interface{}) {
	if !l.__check__(ALL) {
		return
	}
	l.__printf__(format, v...)
}
func (l *Logger) Println(v ...interface{}) {
	if !l.__check__(ALL) {
		return
	}
	l.__println__(v...)
}
func (l *Logger) Panic(v ...interface{}) {
	if !l.__check__(OFF) {
		return
	}
	l.__setPrefix__(l.__level__(OFF))
	l.__panic__(v...)
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	if !l.__check__(OFF) {
		return
	}
	l.__setPrefix__(l.__level__(OFF))
	l.__panicf__(format, v...)
}
func (l *Logger) Panicln(v ...interface{}) {
	if !l.__check__(OFF) {
		return
	}
	l.__setPrefix__(l.__level__(OFF))
	l.__panicln__(v...)
}
func (l *Logger) Fatal(v ...interface{}) {
	if !l.__check__(FATAL) {
		return
	}
	l.__setPrefix__(l.__level__(FATAL))
	l.__fatal__(v...)
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if !l.__check__(FATAL) {
		return
	}
	l.__setPrefix__(l.__level__(FATAL))
	l.__fatalf__(format, v...)
}
func (l *Logger) Fatalln(v ...interface{}) {
	if !l.__check__(FATAL) {
		return
	}
	l.__setPrefix__(l.__level__(FATAL))
	l.__fatalln__(v...)
}
func (l *Logger) Error(v ...interface{}) {
	if !l.__check__(ERROR) {
		return
	}
	l.__setPrefix__(l.__level__(ERROR))
	l.__print__( v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	if !l.__check__(ERROR) {
		return
	}
	l.__setPrefix__(l.__level__(ERROR))
	l.__printf__(format, v...)
}
func (l *Logger) Errorln(v ...interface{}) {
	if !l.__check__(ERROR) {
		return
	}
	l.__setPrefix__(l.__level__(ERROR))
	l.__println__(v...)
}
func (l *Logger) Warn(v ...interface{}) {
	if !l.__check__(WARN) {
		return
	}
	l.__setPrefix__(l.__level__(WARN))
	l.__print__(v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	if !l.__check__(WARN) {
		return
	}
	l.__setPrefix__(l.__level__(WARN))
	l.__printf__(format, v...)
}
func (l *Logger) Warnln(v ...interface{}) {
	if !l.__check__(WARN) {
		return
	}
	l.__setPrefix__(l.__level__(WARN))
	l.__println__(v...)
}
func (l *Logger) Info(v ...interface{}) {
	if !l.__check__(INFO) {
		return
	}
	l.__setPrefix__(l.__level__(INFO))
	l.__print__(v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	if !l.__check__(INFO) {
		return
	}
	l.__setPrefix__(l.__level__(INFO))
	l.__printf__(format, v...)
}
func (l *Logger) Infoln(v ...interface{}) {
	if !l.__check__(INFO) {
		return
	}
	l.__setPrefix__(l.__level__(INFO))
	l.__println__(v...)
}
func (l *Logger) Debug(v ...interface{}) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__setPrefix__(l.__level__(DEBUG))
	l.__print__(v...)
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__setPrefix__(l.__level__(DEBUG))
	l.__printf__(format, v...)
}
func (l *Logger) Debugln(v ...interface{}) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__setPrefix__(l.__level__(DEBUG))
	l.__println__(v...)
}

func Flags() int {
	return __flags__()
}
func SetFlags(flags int) {
	__setFlags__(flags)
}
func Prefix() string {
	return __prefix__()
}
func SetPrefix(prefix string) {
	__setPrefix__(prefix)
}

func Print(v ...interface{}) {
	__level__(ALL)
	__print__(v...)
}
func Printf(format string, v ...interface{}) {
	__level__(ALL)
	__printf__(format, v...)
}
func Println(v ...interface{}) {
	__level__(ALL)
	__println__(v...)
}
func Panic(v ...interface{}) {
	__level__(OFF)
	__panic__(v...)
}
func Panicf(format string, v ...interface{}) {
	__level__(OFF)
	__panicf__(format, v...)
}
func Panicln(v ...interface{}) {
	__level__(OFF)
	__panicln__(v...)
}
func Fatal(v ...interface{}) {
	__level__(FATAL)
	__fatal__(v...)
}
func Fatalf(format string, v ...interface{}) {
	__level__(FATAL)
	__fatalf__(format, v...)
}
func Fatalln(v ...interface{}) {
	__level__(FATAL)
	__fatalln__(v...)
}
func Error(v ...interface{}) {
	__level__(ERROR)
	__print__(v...)
}
func Errorf(format string, v ...interface{}) {
	__level__(ERROR)
	__printf__(format, v...)
}
func Errorln(v ...interface{}) {
	__level__(ERROR)
	__println__(v...)
}
func Warn(v ...interface{}) {
	__level__(WARN)
	__print__(v...)
}
func Warnf(format string, v ...interface{}) {
	__level__(WARN)
	__printf__(format, v...)
}
func Warnln(v ...interface{}) {
	__level__(WARN)
	__println__(v...)
}
func Info(v ...interface{}) {
	__level__(INFO)
	__print__(v...)
}
func Infof(format string, v ...interface{}) {
	__level__(INFO)
	__printf__(format, v...)
}
func Infoln(v ...interface{}) {
	__level__(INFO)
	__println__(v...)
}
func Debug(v ...interface{}) {
	__level__(DEBUG)
	__print__(v...)
}
func Debugf(format string, v ...interface{}) {
	__level__(DEBUG)
	__printf__(format, v...)
}
func Debugln(v ...interface{}) {
	__level__(DEBUG)
	__println__(v...)
}

func Output(calldepth int, s string) error {
	return __output__(calldepth, s)
}
