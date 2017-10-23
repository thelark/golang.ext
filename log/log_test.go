package log

import (
	"testing"
	"os"
	"io/ioutil"
	"strings"
)

var fileName = "test.log"

func TestNew(t *testing.T) {
	fi, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
}
func TestLogger_Fatal(t *testing.T) {
	fi, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		t.Error(err)
	}
	defer fi.Close()
	l := New(fi, "", Ltime|Ldate|Lshortfile)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	msg := "fatal info ~"
	l.Fatal(msg)

	str, err := readFile("test.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, msg) {
		t.Error("错误: 未写入文件")
	}
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
