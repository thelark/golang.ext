package log

import (
	"io"
	"log"
	"fmt"
)

func init() {
	log.SetFlags(log.LstdFlags)
}

func New(writer io.Writer, prefix string, flags int, flagstmpl string) *Logger {
	return __init__(writer, prefix, flags, flagstmpl)
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
func (l *Logger) Flags() int {
	return l.__flags__()
}
func (l *Logger) SetPrefix(prefix string) {
	l.__setPrefix__(prefix)
}
func (l *Logger) Prefix() string {
	return l.__prefix__()
}
func (l *Logger) SetOutput(writer io.Writer) {
	l.__setOutput__(writer)
}
func (l *Logger) Output(calldepth int, s string) error {
	return l.__output__(calldepth, s)
}
func (l *Logger) SetLevel(level __LogLevel__) {
	l.__setLevel__(level)
}
func (l *Logger) Level() __LogLevel__ {
	return l.__level__()
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
	l.__print__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ALL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Printf(format string, v ...interface{}) {
	if !l.__check__(ALL) {
		return
	}
	l.__printf__("%s%s%s%s", l.__level_print__(ALL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprintf(format, v...))
}
func (l *Logger) Println(v ...interface{}) {
	if !l.__check__(ALL) {
		return
	}
	l.__println__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ALL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Panic(v ...interface{}) {
	if !l.__check__(OFF) {
		return
	}
	l.__panic__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ALL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Panicf(format string, v ...interface{}) {
	if !l.__check__(OFF) {
		return
	}
	l.__panicf__("%s%s%s%s", l.__level_print__(ALL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprintf(format, v...))
}
func (l *Logger) Panicln(v ...interface{}) {
	if !l.__check__(OFF) {
		return
	}
	l.__panicln__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ALL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Fatal(v ...interface{}) {
	if !l.__check__(FATAL) {
		return
	}
	l.__fatal__(fmt.Sprintf("%s%s%s%s", l.__level_print__(FATAL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	if !l.__check__(FATAL) {
		return
	}
	l.__fatalf__("%s%s%s%s", l.__level_print__(FATAL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprintf(format, v...))
}
func (l *Logger) Fatalln(v ...interface{}) {
	if !l.__check__(FATAL) {
		return
	}
	l.__fatalln__(fmt.Sprintf("%s%s%s%s", l.__level_print__(FATAL), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Error(v ...interface{}) {
	if !l.__check__(ERROR) {
		return
	}
	l.__print__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ERROR), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	if !l.__check__(ERROR) {
		return
	}
	l.__printf__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ERROR), l.__prefix_print__(), l.__flags_print__(), format), v...)
}
func (l *Logger) Errorln(v ...interface{}) {
	if !l.__check__(ERROR) {
		return
	}
	l.__println__(fmt.Sprintf("%s%s%s%s", l.__level_print__(ERROR), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Warn(v ...interface{}) {
	if !l.__check__(WARN) {
		return
	}
	l.__print__(fmt.Sprintf("%s%s%s%s", l.__level_print__(WARN), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	if !l.__check__(WARN) {
		return
	}
	l.__printf__("%s%s%s%s", l.__level_print__(WARN), l.__prefix_print__(), l.__flags_print__(), fmt.Sprintf(format, v...))
}
func (l *Logger) Warnln(v ...interface{}) {
	if !l.__check__(WARN) {
		return
	}
	l.__println__(fmt.Sprintf("%s%s%s%s", l.__level_print__(WARN), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Info(v ...interface{}) {
	if !l.__check__(INFO) {
		return
	}
	l.__print__(fmt.Sprintf("%s%s%s%s", l.__level_print__(INFO), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Infof(format string, v ...interface{}) {
	if !l.__check__(INFO) {
		return
	}
	l.__printf__("%s%s%s%s", l.__level_print__(INFO), l.__prefix_print__(), l.__flags_print__(), fmt.Sprintf(format, v...))
}
func (l *Logger) Infoln(v ...interface{}) {
	if !l.__check__(INFO) {
		return
	}
	l.__println__(fmt.Sprintf("%s%s%s%s", l.__level_print__(INFO), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Debug(v ...interface{}) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__print__(fmt.Sprintf("%s%s%s%s", l.__level_print__(DEBUG), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
func (l *Logger) Debugf(format string, v ...interface{}) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__printf__("%s%s%s%s", l.__level_print__(DEBUG), l.__prefix_print__(), l.__flags_print__(), fmt.Sprintf(format, v...))
}
func (l *Logger) Debugln(v ...interface{}) {
	if !l.__check__(DEBUG) {
		return
	}
	l.__println__(fmt.Sprintf("%s%s%s%s", l.__level_print__(DEBUG), l.__prefix_print__(), l.__flags_print__(), fmt.Sprint(v...)))
}
