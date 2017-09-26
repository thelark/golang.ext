package file

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * TODO: 文件操作扩展
 */

// TODO: 获取文件路径的相对路径 (相对GOPATH)
func GetFileRelPath(filePath string) string {
	relPath, err := filepath.Rel(fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src"), filePath)
	_ = err
	return relPath
}
// TODO: 判断路径是否存在
func IsExist(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	_ = err
	return false
}
