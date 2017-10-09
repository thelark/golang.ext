package file

import (
	"os"
	"thelark.cn/golang.ext/wrong"
)

/**
 * TODO: 文件操作扩展
 */

// TODO: 判断路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	wrong.Println(err)
	return false
}

// TODO: 创建目录
func MkDIR(path string) {
	if IsExist(path) {
		return
	}
	err := os.MkdirAll(path, 0777)
	wrong.Println(err)
}
