package log

import (
	"io"
	"log"
	"thelark.cn/golang.ext/color/cmd"
	"os"
	"thelark.cn/golang.ext/wrong"
)

type Logger struct {
	level  LogLevel // 日志级别
	output io.Writer
	flag   bool // 是否开启记录日志
	*log.Logger
}

const (
	Ltime         = log.Ltime
	Ldate         = log.Ldate
	Lmicroseconds = log.Lmicroseconds
	Llongfile     = log.Llongfile
	Lshortfile    = log.Llongfile
	LUTC          = log.LUTC
)

func New(fileName string) *Logger {
	fi, err := os.Create(fileName)
	wrong.Println(err)
	return &Logger{level: DEBUG, Logger: log.New(fi, "", log.Ltime), output: fi}
}

// 开启记录日志
func (l *Logger) Open() {
	l.flag = true
}

// 关闭记录日志
func (l *Logger) Close() {
	l.flag = false
}
func (l *Logger) SetFlags(flags uint) {
	if l.Logger != nil {
		l.Logger.SetFlags(int(flags))
	} else {
		log.SetFlags(int(flags))
	}
}
func (l *Logger) SetFile(fileName string) {
	fi, err := os.Create(fileName)
	wrong.Println(err)
	l.Logger.SetOutput(fi)
	l.output = fi
}

func (l *Logger) SetLevel(level LogLevel) {
	l.level = level
}

//FATAL => Magenta // 致命
//ERROR => Red // 错误
//WARN => Yellow // 警告
//INFO=> Cyan // 信息
//DEBUG => Green // 调试
//ALL => White // 所有
func (l *Logger) FATAL(info string) {
	if !l.flag {
		return
	}
	if l.level < FATAL {
		return
	}
	cmd.Magenta().Printf("[FATAL]\t")
	if l.Logger != nil {
		l.Logger.Fatalln(info)
	} else {
		log.Fatalln(info)
	}
}
func (l *Logger) ERROR(info string) {
	if !l.flag {
		return
	}
	if l.level < ERROR {
		return
	}
	cmd.Red().Printf("[ERROR]\t")
	if l.Logger != nil {
		l.Logger.Println(info)
	} else {
		log.Println(info)
	}
}
func (l *Logger) WARN(info string) {
	if !l.flag {
		return
	}
	if l.level < WARN {
		return
	}
	cmd.Yellow().Printf("[WARN]\t")
	if l.Logger != nil {
		l.Logger.Println(info)
	} else {
		log.Println(info)
	}
}
func (l *Logger) INFO(info string) {
	if !l.flag {
		return
	}
	if l.level < INFO {
		return
	}
	cmd.Cyan().Printf("[INFO]\t")
	if l.Logger != nil {
		l.Logger.Println(info)
	} else {
		log.Println(info)
	}
}
func (l *Logger) DEBUG(info string) {
	if !l.flag {
		return
	}
	if l.level < DEBUG {
		return
	}
	cmd.Green().Printf("[DEBUG]\t")
	if l.Logger != nil {
		l.Logger.Println(info)
	} else {
		log.Println(info)
	}
}
