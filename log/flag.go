package log

import "log"

const (
	Ltime         = log.Ltime         // the time: 01:23:23   形如 01:23:23   的时间
	Ldate         = log.Ldate         // the date: 2009/01/23 形如 2009/01/23 的日期
	Lmicroseconds = log.Lmicroseconds // microsecond resolution: 01:23:23.123123.  形如01:23:23.123123   的时间
	Lfile         = Ffile             // full file name and line number: /a/b/c/d.go:23 全路径文件名和行号
	LUTC          = log.LUTC
	LstdFlags     = log.LstdFlags
)

const (
	FNil = iota
)
const (
	Fnanosecond = 1 << iota // 1
	Fsecond                 // 2
	Fminute                 // 4
	Fhour                   // 8
	Fday                    // 16
	Fmonth                  // 32
	Fyear                   // 64
	Ffile                   // 128
)
const (
	Fdate     = Fday | Fmonth | Fyear
	Ftime     = Fsecond | Fminute | Fhour
	FstdFlags = Ftime | Fdate
)
