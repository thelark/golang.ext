package path

import (
	"path/filepath"
	"fmt"
	"os"
	"thelark.cn/golang.ext/wrong"
)

/**
 * TODO: 路径操作扩展
 */

// TODO: 根据文件绝对路径获取文件相对路径 (相对 RefPath)
func GetRelPathFromAbsPath(filePath, RefPath string) string {
	relPath, err := filepath.Rel(RefPath, filePath)
	wrong.Println(err)
	return relPath
}

// TODO: 根据文件据对路径获取文件相对 环境变量中 GOPATH 的路径
func GetRelGoPathFromAbsPath(filePath string) string {
	return GetRelPathFromAbsPath(filePath, fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src"))
}
