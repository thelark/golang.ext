package path

import (
	"path/filepath"
	"fmt"
	"os"
)

/**
 * TODO: 路径操作扩展
 */

// TODO: 获取文件路径的相对路径 (相对GOPATH)
func GetFileRelPath(filePath string) string {
	relPath, err := filepath.Rel(fmt.Sprintf("%s/%s", os.Getenv("GOPATH"), "src"), filePath)
	_ = err
	return relPath
}
