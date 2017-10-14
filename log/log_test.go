package log

import (
	"testing"
	"os"
	"io/ioutil"
	"strings"
)

var fileName = "test.log"

func TestNew(t *testing.T) {
	l := New(fileName)
	if l == nil {
		t.Error("初始化有误！")
	}
}
func TestLogger_FATAL(t *testing.T) {
	l := New(fileName)
	if l == nil {
		t.Error("初始化有误！")
	}
	l.Open()
	l.SetFlags(Ltime | Ldate)
	l.SetLevel(ALL)
	l.FATAL("fatal info ~")
	str, err := readFile("test.log")
	if err != nil {
		t.Error("错误: ", err)
	}
	if !strings.Contains(str, "fatal info ~") {
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
