package log

import "log"

func init() {
	log.SetFlags(log.LstdFlags)
}

func New(fileName string) *Logger {
	return __init__(fileName)
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
func (l *Logger) SetFile(fileName string) {
	l.__setFile__(fileName)
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
func (l *Logger) Fatal(info string) {
	if !l.__check__(FATAL) {
		return
	}
	l.__println__(FATAL, info)
}
func (l *Logger) Error(info string) {
	if !l.__check__(ERROR) {
		return
	}
	l.__println__(ERROR, info)
}
func (l *Logger) Warn(info string) {
	if !l.__check__(WARN) {
		return
	}
	l.__println__(WARN, info)
}
func (l *Logger) Info(info string) {
	if !l.__check__(INFO) {
		return
	}
	l.__println__(INFO, info)
}
func (l *Logger) Debug(info string) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__println__(DEBUG, info)
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
	__setPrefix(prefix)
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
