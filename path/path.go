package path

import (
	"fmt"
	"os"
	"path/filepath"
)

/**
 * TODO: 路径操作扩展
 */

// TODO: 根据文件绝对路径获取文件相对路径 (相对 RefPath)
// @param filePath 文件绝对路径
// @param RefPath 相对目标路径
func GetRelPathFromAbsPath(filePath, RefPath string) (string, error) {
	return filepath.Rel(RefPath, filePath)
}

// TODO: 根据文件据对路径获取文件相对 环境变量中 GOPATH 的路径
// @param filePath 文件绝对路径
func GetRelGoPathFromAbsPath(filePath string) (string, error) {
	return GetRelPathFromAbsPath(filePath, fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src"))
}
