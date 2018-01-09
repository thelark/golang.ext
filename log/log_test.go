package log

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

var fileName = "example.log"

func TestNew(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
}

/*func TestLogger_Fatal(t *testing.T) {
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "fatal info ~"
	l.Fatal(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Fatalf(t *testing.T) {
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "fatalf info ~"
	l.Fatalf(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Fatalln(t *testing.T) {
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "fatalln info ~"
	l.Fatalln(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}

func TestLogger_Panic(t *testing.T) {
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "panic info ~"
	l.Panic(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Panicf(t *testing.T) {
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "panicf info ~"
	l.Panicf(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Panicln(t *testing.T) {
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "panicln info ~"
	l.Panicln(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}*/

func TestLogger_Error(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "error info ~"
	l.Error(msg)
	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Errorf(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "errorf info ~"
	l.Errorf(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Errorln(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "errorln info ~"
	l.Errorln(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}

func TestLogger_Warn(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "warn info ~"
	l.Warn(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Warnf(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "warnf info ~"
	l.Warnf(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Warnln(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "warnln info ~"
	l.Warnln(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}

func TestLogger_Info(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "info info ~"
	l.Info(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Infof(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "infof info ~"
	l.Infof(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Infoln(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "infoln info ~"
	l.Infoln(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}

func TestLogger_Debug(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "debug info ~"
	l.Debug(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Debugf(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "debugf info ~"
	l.Debugf(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Debugln(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "debugln info ~"
	l.Debugln(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}

func TestLogger_Print(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "print info ~"
	l.Print(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Printf(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "printf info ~"
	l.Printf(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}
func TestLogger_Println(t *testing.T) {
	fi, err := openFile()
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ftime|Fdate|Ffile, "")
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetPrefix(" PREFIX ")
	//l.SetFlags(Ftime | Fdate)
	l.SetLevel(ALL)
	msg := "println info ~"
	l.Println(msg)

	str, err := readFile("example.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
}

// 打开文件
func openFile() (*os.File, error) {
	return os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0777)
}

// 读文件
func readFile(path string) (string, error) {
	fi, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	if err != nil {
		return "", err
	}
	return string(fd), nil
}
