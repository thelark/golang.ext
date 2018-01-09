package log

import (
	"log"
	"os"
	"runtime"
	"fmt"
)

var (
	__std__   = log.New(os.Stdout, "", log.LstdFlags)
	__Fatal__ = log.New(os.Stderr, "", log.LstdFlags)
	__Error__ = log.New(os.Stderr, "", log.LstdFlags)
	__Warn__  = log.New(os.Stderr, "", log.LstdFlags)
	__Info__  = log.New(os.Stdout, "", log.LstdFlags)
	__Debug__ = log.New(os.Stdout, "", log.LstdFlags)
	__All__   = log.New(os.Stdout, "", log.LstdFlags)
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
	//__std__.SetFlags(__std__.Flags() &^ Lfile)
	//__Fatal__.SetFlags(__Fatal__.Flags() &^ Lfile)
	//__Error__.SetFlags(__Error__.Flags() &^ Lfile)
	//__Warn__.SetFlags(__Warn__.Flags() &^ Lfile)
	//__Info__.SetFlags(__Info__.Flags() &^ Lfile)
	//__Debug__.SetFlags(__Debug__.Flags() &^ Lfile)
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
	show := make([]interface{}, 0)
	if __std__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__std__.Print(show...)
}
func Printf(format string, v ...interface{}) {
	if __std__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__std__.Printf(format, v...)
}
func Println(v ...interface{}) {
	show := make([]interface{}, 0)
	if __std__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__std__.Println(show...)
}
func Panic(v ...interface{}) {
	show := make([]interface{}, 0)
	if __std__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__std__.Panic(show...)
}
func Panicf(format string, v ...interface{}) {
	if __std__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__std__.Panicf(format, v...)
}
func Panicln(v ...interface{}) {
	show := make([]interface{}, 0)
	if __std__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__std__.Panicln(show...)
}
func Fatal(v ...interface{}) {
	__level__(FATAL)
	show := make([]interface{}, 0)
	if __Fatal__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Fatal__.Fatal(show...)
}
func Fatalf(format string, v ...interface{}) {
	__level__(FATAL)
	if __Fatal__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__Fatal__.Fatalf(format, v...)
}
func Fatalln(v ...interface{}) {
	__level__(FATAL)
	show := make([]interface{}, 0)
	if __Fatal__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Fatal__.Fatalln(show...)
}
func Error(v ...interface{}) {
	__level__(ERROR)
	show := make([]interface{}, 0)
	if __Error__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Error__.Print(show...)
}
func Errorf(format string, v ...interface{}) {
	__level__(ERROR)
	if __Error__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__Error__.Printf(format, v...)
}
func Errorln(v ...interface{}) {
	__level__(ERROR)
	show := make([]interface{}, 0)
	if __Error__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Error__.Println(show...)
}
func Warn(v ...interface{}) {
	__level__(WARN)
	show := make([]interface{}, 0)
	if __Warn__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Warn__.Print(show...)
}
func Warnf(format string, v ...interface{}) {
	__level__(WARN)
	if __Warn__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__Warn__.Printf(format, v...)
}
func Warnln(v ...interface{}) {
	__level__(WARN)
	show := make([]interface{}, 0)
	if __Warn__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Warn__.Println(show...)
}
func Info(v ...interface{}) {
	__level__(INFO)
	show := make([]interface{}, 0)
	if __Info__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Info__.Print(show...)
}
func Infof(format string, v ...interface{}) {
	__level__(INFO)
	if __Info__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__Info__.Printf(format, v...)
}
func Infoln(v ...interface{}) {
	__level__(INFO)
	show := make([]interface{}, 0)
	if __Info__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Info__.Println(show...)
}
func Debug(v ...interface{}) {
	__level__(DEBUG)
	show := make([]interface{}, 0)
	if __Debug__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Debug__.Print(show...)
}
func Debugf(format string, v ...interface{}) {
	__level__(DEBUG)
	if __Debug__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		format = fmt.Sprintf("%s:%d > ", file, line) + format
	}
	__Debug__.Printf(format, v...)
}
func Debugln(v ...interface{}) {
	__level__(DEBUG)
	show := make([]interface{}, 0)
	if __Debug__.Flags()&Lfile == Lfile {
		_, file, line, _ := runtime.Caller(1)
		file, err := __get_rel_go_path_from_abs_path__(file)
		if err != nil {
			panic(err)
		}
		show = append(show, fmt.Sprintf("%s:%d > ", file, line))
	}
	show = append(show, v...)
	__Debug__.Println(show...)
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
