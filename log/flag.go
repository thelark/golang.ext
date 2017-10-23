package log

import "log"

const (
	Ltime         = log.Ltime         // the time: 01:23:23   形如 01:23:23   的时间
	Ldate         = log.Ldate         // the date: 2009/01/23 形如 2009/01/23 的日期
	Lmicroseconds = log.Lmicroseconds // microsecond resolution: 01:23:23.123123.  形如01:23:23.123123   的时间
	Llongfile     = log.Llongfile     // full file name and line number: /a/b/c/d.go:23 全路径文件名和行号
	Lshortfile    = log.Lshortfile    // final file name element and line number: d.go:23. overrides Llongfile 文件名和行号
	LUTC          = log.LUTC
)

