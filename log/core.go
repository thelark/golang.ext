package log

import (
	"fmt"
	"io"
	"log"
	"time"

	"text/template"
	"bytes"
	"runtime"
)

type Logger struct {
	level     __LogLevel__ // 日志级别
	flag      bool         // 是否开启记录日志
	logger    *log.Logger
	flagstmpl string
	prefix    string
	flags     int
}

/**
 * TODO: 初始化日志文件
 * @params writer 文件流
 * @params prefix 日志前缀
 * @params flags 日志 flag
 * @return *Logger
 */
func __init__(out io.Writer, prefix string, flags int, flagstmpl string) *Logger {
	return &Logger{DEBUG, true, log.New(out, "", 0), flagstmpl, prefix, flags}
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
	return l.flags
}

/**
 * TODO: 设置 flags
 * @params flags flags
 */
func (l *Logger) __setFlags__(flags int) {
	l.flags = flags
}

/**
 * TODO: 获取 前缀
 * @return string
 */
func (l *Logger) __prefix__() string {
	return l.prefix
}

/**
 * TODO: 设置 前缀
 * @params prefix 前缀
 */
func (l *Logger) __setPrefix__(prefix string) {
	l.prefix = prefix
}

/**
 * TODO: 设置 级别
 * @params level 日志级别
 */
func (l *Logger) __setLevel__(level __LogLevel__) {
	l.level = level
}

/**
 * TODO: 获取 级别
 */
func (l *Logger) __level__() __LogLevel__ {
	return l.level
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
func (l *Logger) __level_print__(level __LogLevel__) string {
	return fmt.Sprintf("%-9s", fmt.Sprintf("[%s]", __LogLevelName__[level]))
}
func (l *Logger) __prefix_print__() string {
	return fmt.Sprintf("%s", l.prefix)
}
func (l *Logger) __flags_print__() string {
	if l.flagstmpl == "" {
		l.flagstmpl = "{{if ne .Year 0}}{{.Year}}{{end}}{{if ne .Month 0}}/{{.Month}}{{end}}{{if ne .Day 0}}/{{.Day}}{{end}} {{if ne .Hour 0}}{{.Hour}}{{end}}{{if ne .Minute 0}}:{{.Minute}}{{end}}{{if ne .Second 0}}:{{.Second}}{{end}}{{if ne .Nanosecond 0}}.{{.Nanosecond}}{{end}} {{.File}}"
	}
	t := time.Now()

	type flagsStruct struct {
		Year       int
		Month      int
		Day        int
		Hour       int
		Minute     int
		Second     int
		Nanosecond int
		File       string
	}
	tmp := flagsStruct{}

	if l.flags&(FstdFlags|Fnanosecond) != 0 {
		if l.flags&(Fnanosecond) != 0 {
			tmp.Nanosecond = t.Nanosecond()
		}
		if l.flags&(Fsecond|Fminute|Fhour) != 0 {
			if l.flags&(Fsecond) != 0 {
				tmp.Second = t.Second()
			}
			if l.flags&(Fminute) != 0 {
				tmp.Minute = t.Minute()
			}
			if l.flags&(Fhour) != 0 {
				tmp.Hour = t.Hour()
			}
		}
		if l.flags&(Fday|Fmonth|Fyear) != 0 {
			if l.flags&(Fday) != 0 {
				tmp.Day = t.Day()
			}
			if l.flags&(Fmonth) != 0 {
				tmp.Month = int(t.Month())
			}
			if l.flags&(Fyear) != 0 {
				tmp.Year = t.Year()
			}
		}
	}

	if l.flags&Ffile != 0 {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
			tmp.File = fmt.Sprintf("%s:%d > ", file, line)
		} else {
			refPath, err := __get_rel_go_path_from_abs_path__(file)
			if err != nil {
				panic(err)
			}
			tmp.File = fmt.Sprintf("%s:%d > ", refPath, line)
		}
	}
	flagstmpl, err := template.New("flags").Parse(l.flagstmpl)
	if err != nil {
		panic(err)
	}
	retbuffer := bytes.NewBuffer([]byte{})
	flagstmpl.Execute(retbuffer, tmp)
	return retbuffer.String()
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
